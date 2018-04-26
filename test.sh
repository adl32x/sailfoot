#!/bin/bash

echo "Build binary."
go build -o sf

echo "Start testing."
./tests/test_all.sh
