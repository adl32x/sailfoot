#!/bin/bash

set -e

echo "Build binary."
go build -o sf

echo "Start testing."
echo "-------------------"

./tests/test_all.sh $1
