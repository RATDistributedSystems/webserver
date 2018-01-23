package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RATDistributedSystems/mux"
	"github.com/RATDistributedSystems/webserver/ratwebserver"
	"github.com/mholt/caddy/caddy/caddymain"
)

// Command datatype for allowed user commands. Is exported
type Command struct {
	command             string
	usernameRequired    bool
	stockIDRequired     bool
	stockAmountRequired bool
	values              map[string]string
}

type configuration struct {
	address  string
	port     int
	protocol string
}

// TServerConf Will get Exported
var TServerConf = &configuration{
	address:  "localhost",
	port:     44441,
	protocol: "tcp"}

func (c *configuration) loadConf() {
	f, errIO := ioutil.ReadFile("./tserver.json")
	if errIO != nil {
		fmt.Println("Can't read configuration file")
		return
	}
	var configuration interface{}
	err := json.Unmarshal(f, &configuration)
	if err != nil {
		fmt.Printf("Can't parse json: %s", string(f))
		return
	}

	config := configuration.(map[string]interface{})
	// Set struct values
	c.address = config["address"].(string)
	c.port = int(config["port"].(float64))
	c.protocol = config["protocol"].(string)
	return
}

func (c Command) generateCMDString() (cmd string) {
	var buffer bytes.Buffer
	buffer.WriteString(c.command)
	if c.usernameRequired {
		buffer.WriteString(", " + c.values["username"])
	}
	if c.stockIDRequired {
		buffer.WriteString(", " + c.values["stock"])
	}
	if c.stockAmountRequired {
		buffer.WriteString(", " + c.values["amount"])
	}
	cmd = buffer.String()
	return
}

func main() {
	TServerConf.loadConf()
	mux.HandleFunc("/result", requestHandler)
	caddymain.Run()
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	command, err := getPostInformation(r.PostForm)
	if err != nil {
		ratwebserver.ErrorResponse(w, err.Error())
		return
	}

	if command != nil {
		ratwebserver.SendToTServer(TServerConf.address, TServerConf.port, TServerConf.protocol, command.generateCMDString())
		ratwebserver.SuccessResponse(w)
	}
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
		if notNumeric(StockAmtSplice[0]) {
			return nil, fmt.Errorf("Amount: %s not valid number", StockAmtSplice[0])
		}
		mapValues["amount"] = StockAmtSplice[0]
	}

	return cmd, nil
}

func notNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err != nil
}

func createCommandStruct(c string, uname bool, stock bool, amt bool) *Command {
	return &Command{c, uname, stock, amt, nil}
}

func checkForValidCommand(cmd string) (c *Command, e error) {
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
	default:
		c, e = nil, errors.New("Invalid Command")
	}
	return
}
