package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/a0id/mystery/client"
	"github.com/a0id/mystery/server"
	"github.com/a0id/mystery/types"
)

var (
	serverFlag   = flag.String("server", "", "Start a server")
	clientFlag   = flag.String("client", "", "Start a client")
	generateFlag = flag.Bool("generate", false, "Generate an attempt struct")
	loadFlag     = flag.String("load", "", "Load an attempt")
)

func main() {
	flag.Parse()
	// Handle server
	if *serverFlag != "" {
		err := server.InitServer(*serverFlag)
		if err != nil {
			panic(err)
		}
	} else

	// Handle client
	if *clientFlag != "" {
		params := strings.Split(*clientFlag, ":")
		err := client.InitClient(params[0], params[1])
		if err != nil {
			panic(err)
		}
	} else

	// Generate an attempts
	if *generateFlag {
		reader := bufio.NewReader(os.Stdin)

		// Get the username
		fmt.Print("username ? ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		// Get the pin
		fmt.Print("pin ? ")
		rawPin, _ := reader.ReadString('\n')
		pin, err := strconv.Atoi(strings.TrimSuffix(rawPin, "\n"))
		if err != nil {
			panic(err)
		}

		// Get the payload file
		fmt.Print("payload filename ? ")
		payloadFile, _ := reader.ReadString('\n')
		payloadFile = strings.TrimSuffix(payloadFile, "\n")

		// Get the file to write to
		fmt.Print("export filename ? ")
		exportFile, _ := reader.ReadString('\n')
		exportFile = strings.TrimSuffix(exportFile, "\n")

		// Get the passphrase
		fmt.Print("passphrase ? ")
		passphrase, _ := reader.ReadString('\n')
		passphrase = strings.TrimSuffix(passphrase, "\n")

		payload, err := ioutil.ReadFile(payloadFile)
		if err != nil {
			panic(err)
		}

		attempt, err := types.NewAttempt(username, pin, payload)

		encrypted, err := types.EncryptAttempt(attempt, []byte(passphrase))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(exportFile, encrypted, 0600)
		if err != nil {
			panic(err)
		}

		fmt.Println("exported")
	} else

	// Load an attempt from memory
	if *loadFlag != "" {
		raw, err := ioutil.ReadFile(*loadFlag)
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(os.Stdin)

		// Get the username
		fmt.Print("passphrase ? ")
		passphrase, _ := reader.ReadString('\n')
		passphrase = strings.TrimSuffix(passphrase, "\n")

		attempt, err := types.DecryptAttempt(raw, []byte(passphrase))
		if err != nil {
			panic(err)
		}

		fmt.Println(attempt.String())

	}

}
