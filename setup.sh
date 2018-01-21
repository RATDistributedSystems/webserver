#!/bin/bash

go get github.com/mholt/caddy/caddy
go get github.com/caddyserver/builds
wget https://raw.githubusercontent.com/mholt/caddy/master/caddyhttp/httpserver/plugin.go
python add_plugin.py
mv plugin.go $GOPATH/src/github.com/mholt/caddy/caddyhttp/httpserver/
cd $GOPATH/src/github.com/mholt/caddy/caddy
go run build.go
