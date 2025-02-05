#!/bin/bash

for file in ./examples/*/*/*.go; do
    go run $file > /dev/null
    if [ $? -eq 0 ]; then
        echo "$file succeeded"
    else
        echo "$file failed"
        exit 1
    fi
done;