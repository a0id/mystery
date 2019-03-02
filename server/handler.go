package server

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/xoreo/mystery/common"
)

// Handle - Handle an incomming connection
func Handle(conn net.Conn, validationChan chan []byte) {
	// var buffer []byte
	buffer := make([]byte, common.BuffSize)
	for {
		size, err := conn.Read(buffer)
		fmt.Printf("received: %s\n", buffer)
		if err == nil && size > 0 && size < common.BuffSize {
			fmt.Println("good")
			fmt.Printf("received: %s\n", buffer)
			err = ioutil.WriteFile("received", buffer, 0600)
			if err != nil {
				fmt.Println("writing error")
			}
			validationChan <- buffer
			break
		}
	}

}
