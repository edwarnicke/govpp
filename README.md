[govpp](https://github.com/FDio/govpp/blob/master/README.md) is a go binding for the vpp api

This repo provides the go binding for a version of vpp (currently vpp 20.09)

```go
import "github.com/edwarnicke/govpp/binapi"
```

This repo also provides minimal docker containers for running vpp:

- [vpp](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp)
- [vpp-dbg](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp) - debug image

## What to do if you need a differnt version of vpp:

If you want a version of VPP, you can fork *this* repo and add it yourself.
Because things are very heavily patterned and generated, its really quite easy:

1. Change default value for ARG VPP_VERSION in [Dockerfile](https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L1) to the vpp tag or commit id you want to build.
2. ```go generate ./...```
3. Add a GHCR_TOKEN secret to your fork containing a PAT allowing you to push to your [Github Container Registry](https://docs.github.com/en/free-pro-team@latest/packages/managing-container-images-with-github-container-registry/pushing-and-pulling-docker-images)
4. Push a PR to your fork and merge it.  This will push the vpp and vpp-dbg images to your repo.

## What to do if you need to patch vpp:

If you want to add patches to vpp simply add them to the [patch/](https://github.com/edwarnicke/govpp/blob/main/patch/) directory.

