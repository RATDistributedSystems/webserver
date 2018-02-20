#!/bin/bash

# This file assumes you have already run the other setup file
cd ..
git clone https://github.com/RATDistributedSystems/frontend

# Build our app in docker with CGO disabled (static linking) then copy it back outside
CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o webserver

# Build the image
docker build -t webserver .

# Delete the garbage
rm -rf frontend webserver
