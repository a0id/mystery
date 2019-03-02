package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// InitClient - Start a client
func InitClient(ip, port string) error {
	fmt.Printf("client connecting to %s:%s\n", ip, port)

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	err = sendPayload(conn)
	if err != nil {
		return err
	}
	return nil
}

// sendPayload - Send a payload
func sendPayload(conn net.Conn) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the passphrase
	fmt.Print("payload filename ? ")
	payloadFile, _ := reader.ReadString('\n')
	payloadFile = strings.TrimSuffix(payloadFile, "\n")

	payload, err := ioutil.ReadFile(payloadFile)
	if err != nil {
		return err
	}

	conn.Write(payload)
	fmt.Printf("wrote payload '%s' to server\n", payload)

	return nil
}
