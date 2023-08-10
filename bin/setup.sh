#! /bin/bash

set -e

echo '--> Installing Go...'
asdf install

echo '--> Installing fswatch and direnv...'
brew install fswatch direnv

echo '--> Installing YASS...'
cd yass
go install
cd ..

echo '--> Done!'
