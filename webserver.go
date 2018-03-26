package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/RATDistributedSystems/utilities"
	"github.com/julienschmidt/httprouter"
)

var config = utilities.Load()

func main() {
	addrWS, _ := config.GetListnerDetails("webserver")
	addrTS, _ := config.GetServerDetails("transaction")
	log.Printf("Serving on %s", addrWS)
	log.Printf("HTTP Requests will be passed onto %s", addrTS)

	// Enable handlers
	router := httprouter.New()
	router.GET("/", getURL)
	router.GET("/add", getURL)
	router.GET("/buy", getURL)
	router.GET("/buytrigger", getURL)
	router.GET("/commit", getURL)
	router.GET("/quote", getURL)
	router.GET("/sell", getURL)
	router.GET("/selltrigger", getURL)
	router.GET("/summary", getURL)
	router.POST("/result", requestHandler)
	log.Fatal(http.ListenAndServe(addrWS, router))

}

func requestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%s request for %s Origin: %s", r.Method, r.URL, r.RemoteAddr)
	r.ParseForm()
	command, err := getPostInformation(r.PostForm)
	if err != nil {
		ErrorResponse(w, err.Error())
		return
	}

	if command != nil {
		addr, protocol := config.GetServerDetails("transaction")
		err := sendToTServer(addr, protocol, command.String())
		if err != nil {
			ErrorResponse(w, "Couldn't Process Request. Try again later")
			return
		}
		SuccessResponse(w)
	}
}

// SendToTServer sends items to transaction server
func sendToTServer(addr string, protocol string, msg string) error {
	log.Printf("Sending '%s' to %s", msg, addr)
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Printf("Request failed. Couldn't Connect to server %s...", addr)
		return err
	}
	fmt.Fprint(conn, msg+"\n")
	return nil
}
