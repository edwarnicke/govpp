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
git_cherry_pick refs/changes/12/33312/4 # 33312: sr: fix srv6 definition of behavior associated to a LocalSID | https://gerrit.fd.io/r/c/vpp/+/33312
git_cherry_pick refs/changes/13/34713/3 # 34713: vppinfra: improve & test abstract socket | https://gerrit.fd.io/r/c/vpp/+/34713
git_cherry_pick refs/changes/71/32271/15 # 32271: memif: add support for ns abstract sockets | https://gerrit.fd.io/r/c/vpp/+/32271
git_cherry_pick refs/changes/34/34734/2 # 34734: memif: autogenerate socket_ids | https://gerrit.fd.io/r/c/vpp/+/34734
git_cherry_pick refs/changes/26/34726/1 # 34726: interface: add buffer stats api | https://gerrit.fd.io/r/c/vpp/+/34726
git_cherry_pick refs/changes/52/34852/2 # 34852: memif: fix rx/txqueue RC on connected | https://gerrit.fd.io/r/c/vpp/+/34852
git_cherry_pick refs/changes/09/35209/1 # 35209: cnat: Fix conflicting rsession | https://gerrit.fd.io/r/c/vpp/+/35209
git_cherry_pick refs/changes/39/35439/1 # 35439: ip: fix assert in ip4_ttl_inc | https://gerrit.fd.io/r/c/vpp/+/35439
git_cherry_pick refs/changes/38/35438/1 # 35438: af_packet: fix tx stall by retrying failed sendto | https://gerrit.fd.io/r/c/vpp/+/35438

# --------------- Dedicated plugins ---------------
git_cherry_pick refs/changes/64/33264/7 # 33264: pbl: Port based balancer | https://gerrit.fd.io/r/c/vpp/+/33264
git_cherry_pick refs/changes/88/31588/1 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588
git_cherry_pick refs/changes/83/28083/21 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/25 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# --------------- Dedicated plugins ---------------

# NSM cherry-picks
git_cherry_pick refs/changes/45/36945/2 # 36945: wireguard: fix ipv6 handshake packet | https://gerrit.fd.io/r/c/vpp/+/36945
git_cherry_pick refs/changes/18/37018/1 # 37018: wireguard: fix ipv6 payload_length computation | https://gerrit.fd.io/r/c/vpp/+/37018
git_cherry_pick refs/changes/57/36157/3 # 36157: ikev2: fix tunnel direction | https://gerrit.fd.io/r/c/vpp/+/36157

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
