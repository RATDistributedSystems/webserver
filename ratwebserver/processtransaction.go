package ratwebserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configFile struct {
	address  string
	port     int
	protocol string
}

func connectToTransactionServer() {

}

// LoadConf is a function
func LoadConf() {
	var address string
	var port int
	var protocol string

	conf, errIO := ioutil.ReadFile("tserver.conf")
	if errIO != nil {
		address = "localhost"
		port = 44441
		protocol = "tcp"
	}
	var configuration configFile
	err := json.Unmarshal(conf, &configuration)
	if err != nil {

	}
	fmt.Println(configuration)
	fmt.Printf("Address: %s\nPort: %d\nProtocol: %s", address, port, protocol)
}
