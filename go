#!/usr/bin/env bash

set -e
script_dir=$(cd "$(dirname "$0")" ; pwd -P)

goal_build-dev-container() {
  pushd "${script_dir}" > /dev/null
    docker build -t data-derp .
  popd > /dev/null
}

TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    build-dev-container                    - Builds dev container (tagged: data-derp)
"
  exit 1
fi