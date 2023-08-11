#! /bin/bash

set -e

go install

GOPATH=`go env GOPATH`
PATH="$PATH:$GOPATH/bin"

yass queen
