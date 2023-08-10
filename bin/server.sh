#! /bin/bash

set -e

# Watch for changes in the background
fswatch site | xargs -n1 bin/build.sh &

bin/build.sh

go run yass/server/main.go
