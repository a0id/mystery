package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/xoreo/mystery/common"

	"github.com/xoreo/mystery/client"
	"github.com/xoreo/mystery/server"
	"github.com/xoreo/mystery/types"
)

var (
	serverFlag   = flag.String("server", "", "Start a server\nEx: ./main --server <port>")
	clientFlag   = flag.String("client", "", "Start a client\nEx: ./main --client <ip>:<port>")
	generateFlag = flag.Bool("generate", false, "Generate a secure payload (an attempt struct)")
	loadFlag     = flag.String("load", "", "Load a secure payload (an attempt struct)\nEx: ./main --load <filename>")
)

func main() {
	common.InitCommon()
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
		username = strings.TrimSuffix(username, common.NEWLINE)

		// Get the pin
		fmt.Print("pin ? ")
		rawPin, _ := reader.ReadString('\n')
		pin, err := strconv.Atoi(strings.TrimSuffix(rawPin, common.NEWLINE))
		if err != nil {
			panic(err)
		}

		// Get the payload file
		fmt.Print("payload filename ? ")
		payloadFile, _ := reader.ReadString('\n')
		payloadFile = strings.TrimSuffix(payloadFile, common.NEWLINE)

		// Get the file to write to
		fmt.Print("export filename ? ")
		exportFile, _ := reader.ReadString('\n')
		exportFile = strings.TrimSuffix(exportFile, common.NEWLINE)

		// Get the passphrase
		fmt.Print("passphrase ? ")
		passphrase, _ := reader.ReadString('\n')
		passphrase = strings.TrimSuffix(passphrase, common.NEWLINE)

		payload, err := ioutil.ReadFile(payloadFile)
		if err != nil {
			panic(err)
		}

		attempt, err := types.NewAttempt(username, pin, payload)
		if err != nil {
			panic(err)
		}
		encrypted, err := types.EncryptAttempt(*attempt, []byte(passphrase))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(exportFile+".secure", encrypted, 0600)
		if err != nil {
			panic(err)
		}

		fmt.Println("exported attempt")
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
		passphrase = strings.TrimSuffix(passphrase, common.NEWLINE)

		attempt, err := types.DecryptAttempt(raw, []byte(passphrase))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(*loadFlag, attempt.Bytes(), 0600)
		if err != nil {
			panic(err)
		}

		fmt.Println(attempt.String())

	}

	// decoded, err := base64.StdEncoding.DecodeString("ewogICJ1c2VybmFtZSI6IHsKICAgICJ1c2VybmFtZSI6ICJhMGlkIiwKICAgICJwaW4iOiB7CiAgICAgICJwaW4iOiA5MTY0LAogICAgICAibGVuZ3RoIjogNAogICAgfSwKICAgICJoYXNoIjogIlRTRTIwLzROTHVKSmlzVWs0TXhUOXlIYkN1akV5SE4xWlBkZ1VlOHVjaUU9IgogIH0sCiAgInBheWxvYWQiOiAiZEdocGN5QnBjeUIwYUdVZ2NHRjViRzloWkE9PSIsCiAgIm9yaWdpbiI6ICIxOTIuMTY4LjEuMjUzIiwKICAidGltZXN0YW1wIjogIjIwMTktMDMtMDIgMTA6NTM6MzUuMzE3NjUzIC0wNTAwIEVTVCBtPSswLjAwMTQ3NTU4MiIsCiAgImhhc2giOiAiWW1yOFlVOVVtZkFSN1FPNXk4OXFhOW9tZnV0QUpuQkFkaTJzVnV2a3p4UT0iCn0=")
	// if err != nil {
	// 	panic(err)
	// }

	// // decrypted, err := common.AESDecrypt(
	// // 	decoded,
	// // 	[]byte("matt"),
	// // )

	// decrypted, err := types.AttemptFromBytes(decoded)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(decrypted.String())

}
