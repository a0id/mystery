package server

import (
	"fmt"
	"net"
)

// InitServer - Start a server
func InitServer(port string) error {
	// Server setup
	validationChan := make(chan []byte, 1000)
	fmt.Println("server initiated")
	fd, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	clientCount := -100
	for {
		if clientCount > 5 {
			break
		}
		connection, err := fd.Accept()

		if err == nil {
			// Handle the client
			clientCount++
			fmt.Println("accepted client")
			go Handle(connection, validationChan)

			// Read from the channel & validate
			data := <-validationChan
			fmt.Printf("from chan: %s\n", data)
			valid, err := isValid(data)
			if err != nil {
				return err
			}

			// Respond
			if valid {
				connection.Write([]byte("accepted"))
			} else {
				connection.Write([]byte("denied"))
			}
		}
	}
	return nil
}

// isValid - Validate an attempt
func isValid(input []byte) (bool, error) {
	return true, nil // Placeholder for now
}
