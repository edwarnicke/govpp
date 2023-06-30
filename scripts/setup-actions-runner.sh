#!/bin/bash -x
# shellcheck disable=SC2086

runner_token=$1

mkdir actions-runner && cd actions-runner
curl -o actions-runner-linux-x64-2.305.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.305.0/actions-runner-linux-x64-2.305.0.tar.gz
tar xzf ./actions-runner-linux-x64-2.305.0.tar.gz

export RUNNER_ALLOW_RUNASROOT=1
echo | ./config.sh --url https://github.com/networkservicemesh/govpp --token ${runner_token} || exit 1
./run.sh > f.log 2> f.err &
