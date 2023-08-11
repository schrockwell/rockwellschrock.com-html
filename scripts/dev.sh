#! /bin/bash

set -e

# Watch for changes in the background
fswatch web | xargs -n1 yass queen &

# Run the initial build
yass queen

# Start the server
yass server
