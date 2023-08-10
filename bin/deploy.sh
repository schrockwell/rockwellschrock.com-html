#! /bin/bash

set -e

cd yass
go install
cd ..

GOPATH=`go env GOPATH`
PATH="$PATH:$GOPATH/bin"

yass build
