[govpp](https://github.com/FDio/govpp/blob/master/README.md) is a go binding for the vpp api

This repo provides the go binding for several versions of vpp

VPP 20.09:
```go
import "github.com/edwarnicke/govpp/pkg/v2009/binapi"
```

VPP 20.05:
```go
import "github.com/edwarnicke/govpp/pkg/v2005/binapi"
```
This repo also provides minimal docker containers for running vpp:

- [vpp](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp)
- [vpp-dbg](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp) - debug image

### What to do if you need a version of vpp available here:

- Open an [Issue](https://github.com/edwarnicke/govpp/issues) - Likelihood of adding it is high
- Push a PR to add it

If you want a version of VPP that is not of general issue, you can fork *this* repo and add it yourself.
Because things are very heavily patterned and generated, its really quite easy:

1. Add the version you want to the [gen.go](https://github.com/edwarnicke/govpp/blob/main/gen.go) file
2. ```go generate ./...```
3. Add a Docker Build Step for the version to [.github/workflows/ci.yaml](https://github.com/edwarnicke/govpp/blob/main/.github/workflows/ci.yaml#L46)
4. Add the version to the Docker Push Step in [.github/workflows/ci.yaml](https://github.com/edwarnicke/govpp/blob/main/.github/workflows/ci.yaml#L51)
5. Add a GHCR_TOKEN secret to your fork containing a PAT allowing you to push to your [Github Container Registry](https://docs.github.com/en/free-pro-team@latest/packages/managing-container-images-with-github-container-registry/pushing-and-pulling-docker-images)
