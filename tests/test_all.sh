#!/bin/bash

set -e

# Move to script directory.
cd "${0%/*}"

if [ -n "$1" ]; then
    # Single keyword test    
    ../sf -file keyword_tests/test_$1.txt
else
    # Test all
    echo "Running all tests in keywords/"
    echo "------------------------------"
    for file in keyword_tests/test_*; do
        ../sf -driver chromeheadless -file $file
    done
fi



#../sf -file test1.txt
#../sf -file todo.txt

#../sf -driver chromeheadless -file test1.txt
#../sf -driver chromeheadless -file todo.txt