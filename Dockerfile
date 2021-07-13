ARG VPP_VERSION=v21.06
ARG UBUNTU_VERSION=20.04
ARG GOVPP_VERSION=v0.3.5

FROM ubuntu:${UBUNTU_VERSION} as version
ARG VPP_VERSION
ENV VPP_VERSION ${VPP_VERSION}
CMD echo ${VPP_VERSION}

FROM ubuntu:${UBUNTU_VERSION} as vppbuild
ARG VPP_VERSION
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive TZ=US/Central apt-get install -y git make python3 sudo asciidoc
RUN git clone -b ${VPP_VERSION} https://github.com/FDio/vpp.git
WORKDIR /vpp
#COPY patch/ patch/
#RUN git apply patch/*.patch
RUN DEBIAN_FRONTEND=noninteractive TZ=US/Central UNATTENDED=y make install-dep
RUN make pkg-deb
RUN ./src/scripts/version > /vpp/VPP_VERSION

FROM ubuntu:${UBUNTU_VERSION} as vppinstall
ARG VPP_VERSION
COPY --from=vppbuild /var/lib/apt/lists/* /var/lib/apt/lists/
COPY --from=vppbuild [ "/vpp/build-root/libvppinfra_*_amd64.deb", "/vpp/build-root/vpp_*_amd64.deb", "/vpp/build-root/vpp-plugin-core_*_amd64.deb", "/vpp/build-root/vpp-plugin-dpdk_*_amd64.deb", "/pkg/"]
ARG VPP_VERSION
RUN VPP_INSTALL_SKIP_SYSCTL=false apt install -f -y --no-install-recommends /pkg/*.deb ca-certificates iputils-ping iproute2 tcpdump; \
    rm -rf /var/lib/apt/lists/*; \
    rm -rf /pkg

FROM ubuntu:${UBUNTU_VERSION} as vpp
COPY --from=vppinstall / /

FROM vpp as vpp-dbg
ARG VPP_VERSION
WORKDIR /pkg/
COPY --from=vppbuild ["/vpp/build-root/libvppinfra-dev_*_amd64.deb", "/vpp/build-root/vpp-dbg_*_amd64.deb", "/vpp/build-root/vpp-dev_*_amd64.deb", "./" ]
RUN VPP_INSTALL_SKIP_SYSCTL=false apt install -f -y --no-install-recommends ./*.deb


FROM golang:1.15.3-alpine3.12 as binapi-generator
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOBIN=/bin
ARG GOVPP_VERSION
RUN go get git.fd.io/govpp.git/cmd/binapi-generator@${GOVPP_VERSION}

FROM alpine:3.12 as gen
COPY --from=vpp /usr/share/vpp/api/ /usr/share/vpp/api/
COPY --from=binapi-generator /bin/binapi-generator /bin/binapi-generator
COPY --from=vppbuild /vpp/VPP_VERSION /VPP_VERSION
WORKDIR /gen
CMD VPP_VERSION=$(cat /VPP_VERSION) binapi-generator ${PKGPREFIX+-import-prefix ${PKGPREFIX}}






