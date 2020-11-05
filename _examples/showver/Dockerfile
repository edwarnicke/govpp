FROM ghcr.io/edwarnicke/govpp/vpp:20.09 as go
COPY --from=golang:1.15.3-buster /usr/local/go/ /go
ENV PATH ${PATH}:/go/bin
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOBIN=/bin
RUN rm -r /etc/vpp
RUN go get github.com/go-delve/delve/cmd/dlv@v1.5.0

FROM go as build
WORKDIR /build
COPY go.mod go.sum ./
COPY . .
RUN go build -o /bin/showver .

FROM ghcr.io/edwarnicke/govpp/vpp:20.09 as runtime
COPY --from=build /bin/showver /bin/showver
CMD /bin/showver

