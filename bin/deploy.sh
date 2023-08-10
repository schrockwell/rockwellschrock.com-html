#! /bin/bash

set -e

cd yass
go install
cd ..

yass build
