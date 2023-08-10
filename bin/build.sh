#! /bin/bash

which yass > /dev/null

if [ $? -eq 0 ]; then
  yass
else
  go run yass/main.go
fi
