#! /bin/bash

set -e

# Watch for changes in the background
fswatch web | xargs -n1 sh -c 'yass convert && yass queen;' &

# Run the initial build
yass queen

# Start the server
yass server
