#!/usr/bin/env bash

version=$1

mkdir -p results/$version

for dir in deprecate retract mod-dep; do
    pushd clients/$dir
    go mod tidy
    go run . 2>&1 | tee ../../results/$version/$dir.log
    popd
done