#!/bin/bash -x
# shellcheck disable=SC2086

runner_token=$1

cd actions-runner
export RUNNER_ALLOW_RUNASROOT=1
./config.sh remove --token ${runner_token}
