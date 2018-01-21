package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/RATDistributedSystems/mux"
	"github.com/mholt/caddy/caddy/caddymain"
)

type FormInput url.Values

func main() {
	mux.HandleFunc("/add", AddHandler)
	mux.HandleFunc("/buy", BuyHandler)
	mux.HandleFunc("/sell", SellHandler)
	mux.HandleFunc("/quote", QuoteHandler)
	mux.HandleFunc("/summary", SummaryHandler)
	mux.HandleFunc("/buytrigger", BuyTriggerHandler)
	mux.HandleFunc("/selltrigger", SellTriggerHandler)
	mux.HandleFunc("/canceltrigger", CancelTriggerHandler)
	caddymain.Run()
}

func GetUsername(f FormInput) string {
	username := f["username"]
	if username != nil {

	}

}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}

func BuyHandler(w http.ResponseWriter, r *http.Request) {

}

func SellHandler(w http.ResponseWriter, r *http.Request) {

}

func QuoteHandler(w http.ResponseWriter, r *http.Request) {

}

func SummaryHandler(w http.ResponseWriter, r *http.Request) {

}

func BuyTriggerHandler(w http.ResponseWriter, r *http.Request) {

}

func SellTriggerHandler(w http.ResponseWriter, r *http.Request) {

}

func CancelTriggerHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uname_form := r.PostForm["uname"]
	if uname_form != nil {
		fmt.Println(uname_form[0])
	}

}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("test handler success"))
}
