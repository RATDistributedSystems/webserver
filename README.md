# The Front-End Webserver

We chose to go with [Caddy](github.com/mholt/caddy), a webserver built entirely in Go so that our entire stack remains in Golang.

## How to Setup

```
go get github.com/RATDistributedSystems/webserver

# We will get an error but that is OK
# we just want to download it into the right place

cd $GOHOME/github.com/RATDistributedSystems/webserver/setup
./setup.sh
```

Our Caddy handler should now be built

## Use with Docker

Coming soon

## Why isn't stock Caddy sufficient?

Caddy handles all the session management, TLS and all boring stuff, this repo is really just a HTTP handler which serves a few URIs. To do this, this requires us to modify some of the Caddy source code to add in our HTTP middleware.