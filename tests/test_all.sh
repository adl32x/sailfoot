#!/bin/bash

# Move to script directory.
cd "${0%/*}"

../sf -driver chromeheadless -file test1.txt
../sf -driver chromeheadless -file todo.txt