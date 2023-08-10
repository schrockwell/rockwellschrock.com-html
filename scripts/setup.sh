#! /bin/bash

set -e

echo '--> Installing Go...'
asdf install

echo '--> Installing fswatch and direnv...'
brew install fswatch direnv

echo '--> Installing YASS...'
go install

echo '--> Done!'
