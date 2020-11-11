[govpp](https://github.com/FDio/govpp/blob/master/README.md) is a go binding for the vpp api

This repo provides the go binding for a version of vpp (currently vpp 20.09)

```go
import "github.com/edwarnicke/govpp/binapi"
```

This repo also provides minimal docker containers for running vpp:

- [vpp](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp)
- [vpp-dbg](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp) - debug image

### What to do if you need a version of vpp available here:

If you want a version of VPP that is not of general issue, you can fork *this* repo and add it yourself.
Because things are very heavily patterned and generated, its really quite easy:

1. Change [gen.go](https://github.com/edwarnicke/govpp/blob/main/gen.go) to utilize the version you want
2. ```go generate ./...```
3. Change the Docker Build Step to your desired vpp version [.github/workflows/ci.yaml](https://github.com/edwarnicke/govpp/blob/main/.github/workflows/ci.yaml#L46)
4. Change the the Docker Push Step to your desired vpp version [.github/workflows/ci.yaml](https://github.com/edwarnicke/govpp/blob/main/.github/workflows/ci.yaml#L51)
5. Add a GHCR_TOKEN secret to your fork containing a PAT allowing you to push to your [Github Container Registry](https://docs.github.com/en/free-pro-team@latest/packages/managing-container-images-with-github-container-registry/pushing-and-pulling-docker-images)
