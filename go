#!/usr/bin/env bash

set -e
script_dir=$(cd "$(dirname "$0")" ; pwd -P)

goal_build() {
  pushd "${script_dir}" > /dev/null
    docker build -t data-derp/dev-container .
  popd > /dev/null
}

TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    build                    - Builds dev container (tagged: data-derp/dev-container)
"
  exit 1
fi