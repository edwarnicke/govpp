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

# NSM cherry picks
git_cherry_pick refs/changes/03/39503/1 # 39503: vppinfra: fix setns typo | https://gerrit.fd.io/r/c/vpp/+/39503
git_cherry_pick refs/changes/03/39528/6 # 39528: ping: Simple binary API for running ping based on events | https://gerrit.fd.io/r/c/vpp/+/39528

# Calico cherry picks
git_cherry_pick refs/changes/26/34726/3 # 34726: interface: add buffer stats api | https://gerrit.fd.io/r/c/vpp/+/34726

# Copy Calico local patches
git clone -b v3.26.0 https://github.com/projectcalico/vpp-dataplane.git /vpp-dataplane/
cp /vpp-dataplane/vpplink/generated/patches/* patch/

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
