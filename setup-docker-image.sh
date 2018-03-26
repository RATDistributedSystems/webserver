#!/bin/bash

# This file assumes you have already run the other setup file
git submodule update --init

# Build our app in docker with CGO disabled (static linking) then copy it back outside
CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o webserver

# Build the image
docker build -t ratwebserver .

# Delete the garbage
rm -rf frontend webserver
