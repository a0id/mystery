package server

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/xoreo/mystery/common"
	"github.com/xoreo/mystery/types"
)

// InitServer - Start a server
func InitServer(port string) error {
	// Server setup
	validationChan := make(chan []byte)
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
			fmt.Println("client connected")
			go Handle(connection, validationChan)

			// Read from the channel & validate
			data := <-validationChan
			if err != nil {
				return err
			}

			// Check for validity
			valid := isValid(data)

			// Respond
			if valid {
				fmt.Println("accepted client payload")
				connection.Write([]byte("accepted"))
			} else {
				fmt.Println("declined client payload")
				connection.Write([]byte("denied"))
			}
		}
	}
	return nil
}

// isValid - Validate an attempt
func isValid(input []byte) bool {
	// Load passphrase from memory
	passphrase, err := ioutil.ReadFile("passphrase.sec")
	if err != nil {
		return false
	}

	attempt, err := types.DecryptAttempt(input, passphrase)
	if err != nil {
		return false
	}

	// Export the attempt to server memory
	common.CreateDirIfDoesNotExit(common.ExportDir)
	hexHash := fmt.Sprintf("%x", attempt.Hash)[0:8]
	exportFilename := fmt.Sprintf("%s%s.attempt", common.ExportDir, string(hexHash))
	err = ioutil.WriteFile(exportFilename, attempt.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

	return true
}
