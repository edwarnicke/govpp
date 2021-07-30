#!/bin/bash
set -ex

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

VPP_COMMIT=6f6663f3ba814e5f721f7cc679e3098812ad93c3

git_cherry_pick refs/changes/68/32468/3 # 32468: buffers: fix buffer linearization | https://gerrit.fd.io/r/c/vpp/+/32468
git_cherry_pick refs/changes/32/32732/1 # 32732: crypto: fix sw async crypto with chained buffers | https://gerrit.fd.io/r/c/vpp/+/32732
git_cherry_pick refs/changes/33/32833/1 # 32833: ipsec: disable linearization | https://gerrit.fd.io/r/c/vpp/+/32833
git_cherry_pick refs/changes/86/29386/9 # 29386: virtio: DRAFT: multi tx support | https://gerrit.fd.io/r/c/vpp/+/29386
git_cherry_pick refs/changes/21/31321/10 # 31321: devices: add support for pseudo header checksum | https://gerrit.fd.io/r/c/vpp/+/31321

# Revert for now
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/68/32568/6 && git revert --no-edit FETCH_HEAD
git_cherry_pick refs/changes/69/31869/13 # 31869: gso: do not try gro on small packets | https://gerrit.fd.io/r/c/vpp/+/31869
git_cherry_pick refs/changes/76/32476/3 # 32476: gso: handle push flag in gro | https://gerrit.fd.io/r/c/vpp/+/32476

git_cherry_pick refs/changes/82/32482/1 # 32482: virtio: compute cksums in output no offload | https://gerrit.fd.io/r/c/vpp/+/32482
git_cherry_pick refs/changes/83/32483/1 # 32483: virtio: Still init unused txq | https://gerrit.fd.io/r/c/vpp/+/32483

git_cherry_pick refs/changes/71/32871/1 # 32871: devices: Add queues params in create_if | https://gerrit.fd.io/r/c/vpp/+/32871

# IPv6 ND patch (temporary)
git_cherry_pick refs/changes/68/31868/1 # 31868: ip6-nd: silent the source and target checks on given interface | https://gerrit.fd.io/r/c/vpp/+/31868

# --------------- Cnat patches ---------------
git_cherry_pick refs/changes/88/31588/1 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588
# --------------- Cnat patches ---------------

# ------------- Policies patches -------------
git_cherry_pick refs/changes/83/28083/16 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/20 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# ------------- Policies patches -------------

git_cherry_pick refs/changes/35/32235/1  # 32235: dpdk: enable ena interrupt support | https://gerrit.fd.io/r/c/vpp/+/32235

# NSM
git_cherry_pick refs/changes/21/33421/16 # 33421 ping: fix breakage around source address selection (sas)
git_cherry_pick refs/changes/13/32413/6  # 32413 wireguard: move adjacency processing from wireguard_peer to wireguard_interface
git_cherry_pick refs/changes/43/32443/7  # 32443 wireguard: use the same udp-port for multi-tunnel
git_cherry_pick refs/changes/09/32009/5  # 32009 wireguard: add ipv6 support
git_cherry_pick refs/changes/85/32685/4  # 32685 wireguard: add events for peer

if [ "$(ls *.patch 2> /dev/null)" ]; then git apply patch/*.patch;fi
