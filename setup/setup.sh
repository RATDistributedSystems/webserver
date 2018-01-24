#!/bin/bash

# Increase FD imit
# ulimit -n 8192

# Install Deps
go get github.com/mholt/caddy/caddy
go get github.com/caddyserver/builds
go get github.com/RATDistributedSystems/mux

# Get the latest plugin file for us to modify
wget https://raw.githubusercontent.com/mholt/caddy/master/caddyhttp/httpserver/plugin.go
python add_plugin.py
mv plugin.go $GOPATH/src/github.com/mholt/caddy/caddyhttp/httpserver/

# Compile everything
cd ..
run go build server.go

# Now download the frontend
cd ..
git clone https://github.com/RATDistributedSystems/frontend


