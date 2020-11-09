ARG VPP_VERSION=20.09-release
ARG GOVPP_VERSION=v0.3.5
ARG UBUNTU_VERSION=20.04
FROM ubuntu:${UBUNTU_VERSION} as repo
RUN apt-get update && \
     apt-get install -y --no-install-recommends curl ca-certificates debian-archive-keyring gnupg apt-transport-https && \
     (curl -s https://packagecloud.io/install/repositories/fdio/release/script.deb.sh | bash) && \
     (curl -s https://packagecloud.io/install/repositories/fdio/master/script.deb.sh | bash) && \
     apt-get update

FROM ubuntu:${UBUNTU_VERSION} as vpp
COPY --from=repo /var/lib/apt/lists/* /var/lib/apt/lists/
COPY --from=repo /etc/apt/sources.list.d/* /etc/apt/sources.list.d/
ARG VPP_VERSION
RUN apt-get install -y --no-install-recommends ca-certificates iputils-ping iproute2 && \
    VPP_INSTALL_SKIP_SYSCTL=false apt-get install -y --no-install-recommends libvppinfra=${VPP_VERSION} vpp=${VPP_VERSION} vpp-plugin-core=${VPP_VERSION} vpp-plugin-dpdk=${VPP_VERSION}

FROM vpp as vpp-dbg
ARG VPP_VERSION
RUN apt-get install -y --no-install-recommends vpp-dbg=${VPP_VERSION}
WORKDIR /pkg/
RUN chmod 777 . && apt-get download vpp=${VPP_VERSION} vpp-plugin-core=${VPP_VERSION} vpp-plugin-dpdk=${VPP_VERSION} vpp-dev vpp-dbg=${VPP_VERSION}


FROM golang:1.15.3-alpine3.12 as binapi-generator
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOBIN=/bin
ARG GOVPP_VERSION
RUN go get git.fd.io/govpp.git/cmd/binapi-generator@${GOVPP_VERSION}

FROM alpine:3.12 as gen
COPY --from=vpp /usr/share/vpp/api/ /usr/share/vpp/api/
COPY --from=binapi-generator /bin/binapi-generator /bin/binapi-generator
WORKDIR /gen
ARG VPP_VERSION
ENV VPP_VERSION ${VPP_VERSION}
CMD binapi-generator ${PKGPREFIX+-import-prefix ${PKGPREFIX}}






