#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
FILE_NAME="coverage"
OUT_FILE="$FILE_NAME.out"
OUT_HTML="$FILE_NAME.html"
OPEN_FILE="file://$DIR/../$OUT_HTML"

set -e
cd $DIR/..
go get -t ./...
go test -race $(go list ./... | grep -v /examples)  -coverpkg=./... -coverprofile ./$OUT_FILE.tmp
cat $OUT_FILE.tmp | grep -v "mocks" > $OUT_FILE
go tool cover -func ./$OUT_FILE
go tool cover -html $OUT_FILE -o $OUT_HTML
if [[ $1 == "--open" ]]; then
    echo "Opening test coverage file: $OPEN_FILE"
    python3 -mwebbrowser $OPEN_FILE
fi
awk 'BEGIN {exit !('"$(go tool cover -func ./$OUT_FILE | grep 'total:' | awk '{print $3}' | sed -e 's/[%]//g')"' > 70.0)}'