package main

import (
	"log"
	"net/http"

	"github.com/RATDistributedSystems/utilities"
	"github.com/RATDistributedSystems/webserver/ratwebserver"
	"github.com/julienschmidt/httprouter"
)

var serverConfig *utilities.Configuration

func main() {
	serverConfig = utilities.LoadConfigs("config.json")
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
	router.POST("/result", requestHandler)
	log.Fatal(http.ListenAndServe(addrWS, router))

}

func requestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ratwebserver.LogHTTPRequest(r)
	r.ParseForm()
	command, err := ratwebserver.GetPostInformation(r.PostForm)
	if err != nil {
		ratwebserver.ErrorResponse(w, err.Error())
		return
	}

	if command != nil {
		addr, protocol := serverConfig.GetServerDetails("transaction")
		err := ratwebserver.SendToTServer(addr, protocol, command.String())
		if err != nil {
			ratwebserver.ErrorResponse(w, "Couldn't Process Request. Try again later")
			return
		}
		ratwebserver.SuccessResponse(w)
	}
}
