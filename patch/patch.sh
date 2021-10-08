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

# NSM
#git_cherry_pick refs/changes/21/33421/16 # 33421 ping: fix breakage around source address selection (sas)
git_cherry_pick refs/changes/49/33449/9  # 33449 ip: source address selection
git_cherry_pick refs/changes/56/33156/3  # 33156 ip-neighbor: GARP sent to bogus ip address # added to avoid merge conflict with 33495
git_cherry_pick refs/changes/95/33495/2  # 33495 arp: source address selection
git_cherry_pick refs/changes/13/32413/9  # 32413 wireguard: move adjacency processing from wireguard_peer to wireguard_interface
git_cherry_pick refs/changes/43/32443/11  # 32443 wireguard: use the same udp-port for multi-tunnel
git_cherry_pick refs/changes/09/32009/12  # 32009 wireguard: add ipv6 support
git_cherry_pick refs/changes/85/32685/12  # 32685 wireguard: add events for peer
git_cherry_pick refs/changes/93/33993/2  # 33993 wireguard: peers dump fix
git_cherry_pick refs/changes/03/33303/2 # 33303 memif: fix offset
git_cherry_pick refs/changes/20/33020/2 # 33020 l3xc: reset dpo on delete
git_cherry_pick refs/changes/68/33568/1  # 33568 ip: check if interface has link-local address (addition)
git_cherry_pick refs/changes/58/33558/1  # 33558 ip: check if interface has link-local address
git_cherry_pick refs/changes/29/33729/2  # 33729 ping: set fib_index for lookup_node

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
