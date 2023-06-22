[govpp](https://github.com/FDio/govpp/blob/master/README.md) is a go binding for the vpp api

This repo provides the go binding for a version of vpp.

```go
import "github.com/edwarnicke/govpp/binapi"
```

This repo also provides minimal docker containers for running vpp:

- [vpp](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp)
- [vpp-dbg](https://github.com/users/edwarnicke/packages/container/package/govpp%2Fvpp) - debug image

## How to match the version of vpp you want to the go bindings.

Whenever a combination of [VPP_VERSION](https://github.com/edwarnicke/govpp/blob/e0e3b4843cf510bccc6e6a6eac43729a5e5fa42f/Dockerfile#L1)
and [cherrypicks/patches](https://github.com/edwarnicke/govpp/blob/e0e3b4843cf510bccc6e6a6eac43729a5e5fa42f/patch/patch.sh) is used to build vpp/vpp-dbg docker containers,
they will be tagged in the docker repo.  The corresponding tag will be laid on the main branch containing the corresponding
binapi.

So simply run in the same directory as your go.mod file:

```bash
go get github.com/edwarnicke/govpp/binapi@${tag}
```

Where `${tag}` is the tag of the vpp container you are consuming.

For example:

```bash
go get github.com/edwarnicke/govpp/binapi@v21.06.0-9-16f166164
```

To set your go dependency for govpp to match the docker container you are using.

## What to do if you need a different version of vpp:

If you want a version of VPP, you can fork *this* repo and add it yourself.
Because things are very heavily patterned and generated, its really quite easy:

1. Change default value for ARG VPP_VERSION in [Dockerfile](https://github.com/edwarnicke/govpp/blob/main/Dockerfile#L1) to the vpp tag or commit id you want to build.
2. ```go generate ./...```
3. Push a PR to your fork and merge it.  This will push the vpp and vpp-dbg images to your repo.

## What to do if you need to cherrypick into VPP

See [patch/patch.sh].  Simply add your cherrypicks to that file.

## What to do if you need to patch vpp:

If you want to add patches to vpp simply add them with the `.patch` suffix to the [patch/](https://github.com/edwarnicke/govpp/blob/main/patch/) directory.

## What to do if you need a different version of govpp

Change the version of [govpp in the go.mod file](https://github.com/edwarnicke/govpp/blob/main/go.mod#L5)

## How the magic works ##

See [How the Magic Works](https://github.com/edwarnicke/govpp/issues/16).
