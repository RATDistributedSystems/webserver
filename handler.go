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
	m_values := make(map[string]string)

	// Check Command Format in HTTP POST
	f_command := f["command"]
	if f_command == nil || len(f_command) != 1 {
		return nil, errors.New("Missing/Invalid Command Structure")
	}

	// Get All the Parameters out of the POST request
	m_values["command"] = f_command[0]
	cmd, err := checkForValidCommand(m_values["command"])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Command: %s not found", m_values["command"]))
	}
	cmd.values = m_values

	// Get Username if its required
	if cmd.usernameRequired {
		formUsername := f["username"]
		if formUsername == nil || len(formUsername) != 1 || formUsername[0] == "" {
			return nil, errors.New("Missing/Invalid Required Username")
		}
		m_values["username"] = formUsername[0]
	}

	if cmd.stockIDRequired {
		f_stock_id := f["stock"]
		if f_stock_id == nil || len(f_stock_id) != 1 || f_stock_id[0] == "" {
			return nil, errors.New("Missing/Invalid Required Stock ID")
		}
		m_values["stock"] = f_stock_id[0]
	}

	if cmd.stockAmountRequired {
		f_stock_amt := f["amount"]
		if f_stock_amt == nil || len(f_stock_amt) != 1 || f_stock_amt[0] == "" {
			return nil, errors.New("Missing/Invalid Required Stock Amount")
		}
		fmt.Println(f_stock_amt[0])
		m_values["amount"] = f_stock_amt[0]
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

	}
}
