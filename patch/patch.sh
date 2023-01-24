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
git_cherry_pick refs/changes/13/34713/4 # 34713: vppinfra: improve & test abstract socket | https://gerrit.fd.io/r/c/vpp/+/34713
git_cherry_pick refs/changes/71/32271/16 # 32271: memif: add support for ns abstract sockets | https://gerrit.fd.io/r/c/vpp/+/32271
git_cherry_pick refs/changes/34/34734/3 # 34734: memif: autogenerate socket_ids | https://gerrit.fd.io/r/c/vpp/+/34734
git_cherry_pick refs/changes/26/34726/1 # 34726: interface: add buffer stats api | https://gerrit.fd.io/r/c/vpp/+/34726
git_cherry_pick refs/changes/05/35805/2 # 35805: dpdk: add intf tag to dev{} subinput | https://gerrit.fd.io/r/c/vpp/+/35805

# --------------- Dedicated plugins ---------------
git_cherry_pick refs/changes/64/33264/7 # 33264: pbl: Port based balancer | https://gerrit.fd.io/r/c/vpp/+/33264
git_cherry_pick refs/changes/88/31588/4 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588
git_cherry_pick refs/changes/83/28083/21 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/28 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# --------------- Dedicated plugins ---------------

# NSM cherry-picks
git_cherry_pick refs/changes/74/37274/5 # 37274: af_xdp: fix xdp socket create fail | https://gerrit.fd.io/r/c/vpp/+/37274
git_cherry_pick refs/changes/63/37763/4 # 37763: wireguard: add local variable | https://gerrit.fd.io/r/c/vpp/+/37763
git_cherry_pick refs/changes/01/38001/9 # 38001: wireguard: sending the first handshake | https://gerrit.fd.io/r/c/vpp/+/38001
git_cherry_pick refs/changes/00/38000/4 # 38000: wireguard: update ESTABLISHED flag | https://gerrit.fd.io/r/c/vpp/+/38000

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
