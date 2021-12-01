#!/bin/bash
set -ex

git config --global user.name "John Doe"
git config --global user.email "johndoe@example.com"

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

# Calico cherry picks
git_cherry_pick refs/changes/86/29386/9 # 29386: virtio: DRAFT: multi tx support | https://gerrit.fd.io/r/c/vpp/+/29386
git_cherry_pick refs/changes/71/32271/7 # 32271: memif: add support for ns abstract sockets | https://gerrit.fd.io/r/c/vpp/+/32271
git_cherry_pick refs/changes/57/33557/5 # 33557: ip: unlock_fib on if delete | https://gerrit.fd.io/r/c/vpp/+/33557
git_cherry_pick refs/changes/49/34249/2 # 34249: devices: default mode eth in the api | https://gerrit.fd.io/r/c/vpp/+/34249

git_cherry_pick refs/changes/85/34585/1 # 34585: devices: fix af_packet GSO check | https://gerrit.fd.io/r/c/vpp/+/34585

# --------------- Dedicated plugins ---------------
git_cherry_pick refs/changes/64/33264/3 # 33264: pbl: Port based balancer | https://gerrit.fd.io/r/c/vpp/+/33264
git_cherry_pick refs/changes/88/31588/1 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588
git_cherry_pick refs/changes/83/28083/20 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/24 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# --------------- Dedicated plugins ---------------

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
