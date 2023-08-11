#! /bin/bash

set -e

echo '--> Installing Go...'
asdf install

echo '--> Installing Homebrew dependencies...'
brew install fswatch direnv

echo '--> Installing YASS...'
go install

echo '--> Done!'
