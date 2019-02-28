package main

import (
	"fmt"
	"io/ioutil"

	"github.com/a0id/mystery/types"
)

func main() {
	attempt, err := types.NewAttempt(
		"a0id",
		9164,
		[]byte("this is the payload"),
	)
	if err != nil {
		panic(err)
	}

	secureAttempt, err := types.EncryptAttempt(attempt, []byte("password"))
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("secureattempt", secureAttempt, 0600)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote to file")

}
