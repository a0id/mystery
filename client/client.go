package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/xoreo/mystery/common"
)

// InitClient - Start a client
func InitClient(ip, port string) error {
	// Connect to the server
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Write payload to server
	err = sendPayload(conn)
	if err != nil {
		return err
	}

	// Read response
	buffer := make([]byte, common.BuffSize)
	for {
		size, err := conn.Read(buffer)
		if err == nil && size > 0 && size < common.BuffSize {
			fmt.Printf("server: %s\n", buffer[0:size])
			break
		}
	}

	return nil
}

// sendPayload - Send a payload
func sendPayload(conn net.Conn) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the payload
	fmt.Print("payload filename ? ")
	payloadFile, _ := reader.ReadString('\n')
	payloadFile = strings.TrimSuffix(payloadFile, "\n")

	payload, err := ioutil.ReadFile(payloadFile)
	if err != nil {
		return err
	}

	// Writ eto server
	conn.Write(payload)

	return nil
}
