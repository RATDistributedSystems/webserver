package main

import (
	"fmt"
	"net/http"

	"github.com/RATDistributedSystems/mux"
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	mux.HandleFunc("/test", TestHandler)
	mux.HandleFunc("/login", LoginHandler)
	caddymain.Run()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("%+v\n", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("test handler success"))
}
