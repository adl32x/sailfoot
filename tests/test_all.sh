#!/bin/bash

set -e

# Move to script directory.
cd "${0%/*}"

echo $1

if [ -n "$1" ]; then
    # Single keyword test    
    ../sf -file keywords/test_$1.txt
else
    # Test all
    echo "Running all tests in keywords/"
    echo "------------------------------"
    for file in keywords/*; do
        ../sf -file $file
    done
fi



#../sf -file test1.txt
#../sf -file todo.txt

#../sf -driver chromeheadless -file test1.txt
#../sf -driver chromeheadless -file todo.txt