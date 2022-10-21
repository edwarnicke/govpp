ARG VPP_VERSION=v22.10-rc2
ARG UBUNTU_VERSION=20.04
ARG GOVPP_VERSION=v0.3.5

FROM ubuntu:${UBUNTU_VERSION} as vppbuild
ARG VPP_VERSION
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive TZ=US/Central apt-get install -y git make python3 sudo asciidoc
RUN git clone https://github.com/FDio/vpp.git
WORKDIR /vpp
RUN git checkout ${VPP_VERSION}
COPY patch/ patch/
RUN test -x "patch/patch.sh" && ./patch/patch.sh || exit 1
RUN DEBIAN_FRONTEND=noninteractive TZ=US/Central UNATTENDED=y make install-dep
RUN make pkg-deb
RUN ./src/scripts/version > /vpp/VPP_VERSION

FROM vppbuild as version
CMD cat /vpp/VPP_VERSION

FROM ubuntu:${UBUNTU_VERSION} as vppinstall
COPY --from=vppbuild /var/lib/apt/lists/* /var/lib/apt/lists/
COPY --from=vppbuild [ "/vpp/build-root/libvppinfra_*_amd64.deb", "/vpp/build-root/vpp_*_amd64.deb", "/vpp/build-root/vpp-plugin-core_*_amd64.deb", "/vpp/build-root/vpp-plugin-dpdk_*_amd64.deb", "/pkg/"]
RUN VPP_INSTALL_SKIP_SYSCTL=false apt install -f -y --no-install-recommends /pkg/*.deb ca-certificates iputils-ping iproute2 tcpdump iptables; \
    rm -rf /var/lib/apt/lists/*; \
    rm -rf /pkg

FROM ubuntu:${UBUNTU_VERSION} as vpp
COPY --from=vppinstall / /

FROM vpp as vpp-dbg
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
