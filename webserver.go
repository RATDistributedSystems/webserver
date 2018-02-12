package main

import (
	"log"
	"net/http"

	"github.com/RATDistributedSystems/utilities"
	"github.com/RATDistributedSystems/webserver/ratwebserver"
	"github.com/julienschmidt/httprouter"
)

var serverConfig = utilities.GetConfigurationFile("config.json")

func main() {
	addrWS, _ := serverConfig.GetServerDetails("webserver")
	addrTS, _ := serverConfig.GetServerDetails("transaction")
	log.Printf("Serving on %s", addrWS)
	log.Printf("HTTP Requests will be passed onto %s", addrTS)

	// Enable handlers
	router := httprouter.New()
	router.GET("/", ratwebserver.GetURL)
	router.GET("/add", ratwebserver.GetURL)
	router.GET("/buy", ratwebserver.GetURL)
	router.GET("/buytrigger", ratwebserver.GetURL)
	router.GET("/commit", ratwebserver.GetURL)
	router.GET("/quote", ratwebserver.GetURL)
	router.GET("/sell", ratwebserver.GetURL)
	router.GET("/selltrigger", ratwebserver.GetURL)
	router.GET("/summary", ratwebserver.GetURL)
	router.POST("/result", ratwebserver.RequestHandler)
	log.Fatal(http.ListenAndServe(addrWS, router))

}
