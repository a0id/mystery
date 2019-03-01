package server

import (
	"fmt"
	"net"

	"github.com/xoreo/go-basics/networking/handler"
)

// InitServer - Start a server
func InitServer(port string) error {
	fmt.Println("server initiated")
	fd, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	clientCount := 0
	for {
		if clientCount > 5 {
			break
		}
		connection, err := fd.Accept()

		if err == nil {
			clientCount++
			fmt.Println("accepted client.")
			go handler.Handle(connection)
		}
	}
	return nil
}
