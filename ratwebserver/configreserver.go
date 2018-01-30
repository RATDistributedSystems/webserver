package ratwebserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var usedConfiguration Configuration

type jsonConfiguration struct {
	address  string
	port     string
	protocol string
}

type Configuration struct {
	Webserver   *jsonConfiguration
	Transaction *jsonConfiguration
	HTMLLoc     string
}

var (
	webserverConf = &jsonConfiguration{
		address:  "0.0.0.0",
		port:     "44440",
		protocol: "http",
	}

	transactionserverConf = &jsonConfiguration{
		address:  "localhost",
		port:     "44441",
		protocol: "tcp",
	}

	htmlLoc = "../frontend"
)

func LoadConf() *Configuration {
	// setting defaults
	usedConfiguration.Webserver = webserverConf
	usedConfiguration.Transaction = transactionserverConf
	usedConfiguration.HTMLLoc = htmlLoc

	f, errIO := ioutil.ReadFile("./config.json")
	if errIO != nil {
		log.Println("Can't read configuration file. Using defaults provided")
		return &usedConfiguration
	}

	var conf map[string]interface{}
	err := json.Unmarshal(f, &conf)
	if err != nil {
		log.Printf("Can't parse json: %s", string(f))
		return &usedConfiguration
	}

	webserver := conf["webserver"].(map[string]interface{})
	webserverConf.address = webserver["address"].(string)
	webserverConf.port = webserver["port"].(string)

	transaction := conf["transaction"].(map[string]interface{})
	transactionserverConf.address = transaction["address"].(string)
	transactionserverConf.port = transaction["port"].(string)
	transactionserverConf.protocol = transaction["protocol"].(string)

	return &usedConfiguration
}

func (c *Configuration) GetServerAddress() string {
	return c.Webserver.address + ":" + c.Webserver.port
}

func (c *Configuration) GetTransactionAddress() string {
	return c.Transaction.address + ":" + c.Transaction.port
}

func (c *Configuration) GetTransactionProtocol() string {
	return c.Transaction.protocol
}

func (c *Configuration) GetHTMLLocation() string {
	return c.HTMLLoc
}
