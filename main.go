package main

import (
	"flag"
	"strings"

	"github.com/a0id/mystery/client"
	"github.com/a0id/mystery/server"
)

var (
	serverFlag = flag.String("server", "3030", "Start a server")
	clientFlag = flag.String("client", "", "Start a client")
)

func main() {
	// Handle server
	if *serverFlag != "" {
		err := server.InitServer(*serverFlag)
		if err != nil {
			panic(err)
		}
	}

	// Handle client
	if *clientFlag != "" {
		params := strings.Split(*clientFlag, ":")
		err := client.InitClient(params[0], params[1])
		if err != nil {
			panic(err)
		}
	}

}
