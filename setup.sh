#!/bin/bash

go get github.com/mholt/caddy/caddy
go get github/caddyserver/builds
cd $GOPATH/src/github.com/mholt/caddy/caddy
wget https://raw.githubusercontent.com/mholt/caddy/master/caddyhttp/httpserver/plugin.go
python add_plugin.py
mv plugin.go $GOPATH/src/github.com/mholt/caddy/caddyhttp/httpserver/
go run build.go
