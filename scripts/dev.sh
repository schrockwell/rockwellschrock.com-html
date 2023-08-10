#! /bin/bash

set -e

# Watch for changes in the background
fswatch web | xargs -n1 yass build &

# Run the initial build
yass build

# Start the server
yass server
