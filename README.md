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

1. Build docker image yourself

First make sure the `/setup/setup.sh` has been executed at least once. Now from `setup` execute `/setup-docker-image`

This will create the image `ratwebserver` which can be run as follows

`docker run -p 44440:44440 ratwebserver`

2. Use Docker pull

Coming soon ...

## Why isn't stock Caddy sufficient?

Caddy handles all the session management, TLS and all boring stuff, this repo is really just a HTTP handler which serves a few URIs. To do this, this requires us to modify some of the Caddy source code to add in our HTTP middleware.