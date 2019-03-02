package server

import (
	"fmt"
	"net"

	"github.com/xoreo/mystery/common"
)

// Handle - Handle an incomming connection
func Handle(conn net.Conn, validationChan chan []byte) {
	// defer conn.Close()
	buffer := make([]byte, common.BuffSize)
	for {
		size, err := conn.Read(buffer)
		if err == nil && size > 0 && size < common.BuffSize {
			fmt.Printf("\n\n===== BEGIN PAYLOAD =====\n%s\n===== END PAYLOAD =====\n\n", buffer)

			// Write to the channel and close the connection
			validationChan <- buffer[0:size]
			break
		}
	}

}
