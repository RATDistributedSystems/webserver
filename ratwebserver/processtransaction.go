package ratwebserver

import (
	"fmt"
	"net"
)

// SendToTServer sends items to transaction server
func SendToTServer(ip string, port int, protocol string, msg string) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("Attempting to connect to %s...", addr)
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Printf("\tCouldn't Connect to server %s:%d...\n", addr, port)
	}
	fmt.Fprint(conn, msg+"\n")
}
