package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/RATDistributedSystems/mux"
	"github.com/RATDistributedSystems/webserver/ratwebserver"
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	mux.HandleFunc("/result", requestHandler)
	caddymain.Run()
}

// Command datatype for allowed user commands. Is exported
type Command struct {
	command             string
	usernameRequired    bool
	stockIDRequired     bool
	stockAmountRequired bool
	values              map[string]string
}

func createCommandStruct(c string, uname bool, stock bool, amt bool) *Command {
	return &Command{c, uname, stock, amt, nil}
}

func checkForValidCommand(cmd string) (c *Command, e error) {
	// Like a default case
	c, e = nil, errors.New("Invalid Command")

	switch cmd {
	case "ADD":
		c, e = createCommandStruct(cmd, true, false, true), nil
	case "BUY":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "SELL":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "QUOTE":
		c, e = createCommandStruct(cmd, true, true, false), nil
	case "COMMIT_BUY":
		c, e = createCommandStruct(cmd, true, false, false), nil
	case "COMMIT_SELL":
		c, e = createCommandStruct(cmd, true, false, false), nil
	case "CANCEL_BUY":
		c, e = createCommandStruct(cmd, true, false, false), nil
	case "CANCEL_SELL":
		c, e = createCommandStruct(cmd, true, false, false), nil
	case "SET_BUY_AMOUNT":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "SET_BUY_TRIGGER":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "CANCEL_SET_BUY":
		c, e = createCommandStruct(cmd, true, true, false), nil
	case "SET_SELL_AMOUNT":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "SET_SELL_TRIGGER":
		c, e = createCommandStruct(cmd, true, true, true), nil
	case "CANCEL_SET_SELL":
		c, e = createCommandStruct(cmd, true, true, false), nil
	case "DUMPLOG":
		c, e = createCommandStruct(cmd, true, false, false), nil
	case "DISPLAY_SUMMARY":
		c, e = createCommandStruct(cmd, true, false, false), nil
	}
	return
}

func getPostInformation(f url.Values) (*Command, error) {
	mapValues := make(map[string]string)

	// Check Command Format in HTTP POST
	commandSplice := f["command"]
	if commandSplice == nil || len(commandSplice) != 1 {
		return nil, errors.New("Missing/Invalid Command Structure")
	}

	// Get All the Parameters out of the POST request
	mapValues["command"] = commandSplice[0]
	cmd, err := checkForValidCommand(mapValues["command"])
	if err != nil {
		return nil, fmt.Errorf("Command: %s not found", mapValues["command"])
		//errors.New(fmt.Sprintf())
	}
	cmd.values = mapValues

	// Get Username if its required
	if cmd.usernameRequired {
		usernameSplice := f["username"]
		if usernameSplice == nil || len(usernameSplice) != 1 || usernameSplice[0] == "" {
			return nil, errors.New("Missing/Invalid Required Username")
		}
		mapValues["username"] = usernameSplice[0]
	}

	if cmd.stockIDRequired {
		stockIDSplice := f["stock"]
		if stockIDSplice == nil || len(stockIDSplice) != 1 || stockIDSplice[0] == "" {
			return nil, errors.New("Missing/Invalid Required Stock ID")
		}
		mapValues["stock"] = stockIDSplice[0]
	}

	if cmd.stockAmountRequired {
		StockAmtSplice := f["amount"]
		if StockAmtSplice == nil || len(StockAmtSplice) != 1 || StockAmtSplice[0] == "" {
			return nil, errors.New("Missing/Invalid Required Stock Amount")
		}
		mapValues["amount"] = StockAmtSplice[0]
	}

	return cmd, nil
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	command, err := getPostInformation(r.PostForm)
	if err != nil {
		ratwebserver.ErrorResponse(w, err.Error())
		return
	}

	if command != nil {
		//do processing
		ratwebserver.SuccessResponse(w)
	}
}
