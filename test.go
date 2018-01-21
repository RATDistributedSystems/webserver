package main

import (
    "net/http"

    "github.com/mholt/caddy/caddy/caddymain"
    "github.com/RATDistributedSystems/mux"
)

func main() {
    mux.HandleFunc("/test", TestHandler)
    caddymain.Run()
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("test handler success"))
}