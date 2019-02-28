package main

import (
	"fmt"

	"github.com/a0id/mystery/common"
)

func main() {
	// err := common.EncryptRSA([]byte("test text!"), "testwrite")
	// if err != nil {
	// 	panic(err)
	// }

	// err := common.EncryptFile(
	// 	"encrypted.txt",
	// 	[]byte("test"),
	// 	"password",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	decrypted, err := common.DecryptFile(
		"encrypted.txt",
		"password",
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decrypted: %s\n", decrypted)
}
