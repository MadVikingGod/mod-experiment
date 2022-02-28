#!/usr/bin/env bash

version=$1

mkdir -p results/$version

for dir in deprecate retract mod-dep deprecate-updated retract-updated mod-dep-updated; do
    pushd clients/$dir
    go mod tidy | tee ../../results/$version/$dir.log
    go run . 2>&1 | tee -a ../../results/$version/$dir.log
    popd
done