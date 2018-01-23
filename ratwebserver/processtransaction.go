package ratwebserver

import (
	"fmt"
	"net"
)

// SendToTServer sends items to transaction server
func SendToTServer(addr string, port int, protocol string, msg string) {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Printf("Couldn't Connect to server %s:%d...\n", addr, port)
	}
	fmt.Fprint(conn, msg)

}
