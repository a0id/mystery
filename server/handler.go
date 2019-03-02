package server

import (
	"fmt"
	"net"
)

// Handle - Handle an incomming connection
func Handle(conn net.Conn, validationChan chan []byte) {
	var buffer []byte
	for {
		size, err := conn.Read(buffer)
		if err == nil && size > 0 {
			fmt.Printf("received: %s\n", buffer)
			validationChan <- buffer
		}
	}
}
