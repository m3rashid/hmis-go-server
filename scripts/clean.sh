#!/bin/bash
# Path: scripts/clean.sh

declare -a folders=("node_modules" "build" "dist" ".idea")

for i in "${folders[@]}"
do
  find . -name "$i" -type d -prune -exec rm -rf '{}' +
done
