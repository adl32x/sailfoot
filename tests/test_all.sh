#!/bin/bash

# Move to script directory.
cd "${0%/*}"

../sf -file test1.txt
../sf -file todo.txt