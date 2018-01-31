# The Front-End Webserver

We chose to go with [HttpRouter](http://github.com/julienschmidt/httprouter), a webserver router built entirely in Go so that our entire stack remains in Golang.

## How to Setup

```
go get github.com/RATDistributedSystems/webserver
cd $GOHOME/github.com/RATDistributedSystems/webserver
go build server.go
```

Our Caddy handler should now be built

## Use with Docker

1. Build docker image yourself

From `setup` execute `/setup-docker-image.sh`

This will create the image `ratwebserver` which can be run as follows

`docker run -p 44440:44440 ratwebserver`

The first number after `-p` is the "real" port, while the second number is the docker port

2. Use Docker pull

`docker pull asinha94/seng468_webserver:<tag>` where `<tag>` is the date of image creation in `ddmmyyyy`.

This can then be run as 

`docker run -p 44440:44440 asinha94/seng468_webserver:24012018`