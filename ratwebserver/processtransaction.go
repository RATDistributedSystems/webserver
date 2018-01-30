package ratwebserver

import (
	"fmt"
	"log"
	"net"
)

// SendToTServer sends items to transaction server
func SendToTServer(addr string, protocol string, msg string) error {
	log.Printf("Attempting to connect to %s...", addr)
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Printf("Request failed. Couldn't Connect to server %s...", addr)
		return err
	}
	fmt.Fprint(conn, msg+"\n")
	return nil
}
