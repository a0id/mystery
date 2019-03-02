package types

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEncrypt(t *testing.T) {
	attempt, err := NewAttempt(
		"a0id",
		9164,
		[]byte("this is the payload"),
	)
	if err != nil {
		t.Fatal(err)
	}

	secureAttempt, err := EncryptAttempt(*attempt, []byte("password"))
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("../secureattempt", secureAttempt, 0600)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wrote encrypted to file")
}

func TestDecrypt(t *testing.T) {
	rawAttempt, err := ioutil.ReadFile("../secureattempt")
	if err != nil {
		t.Fatal(err)
	}

	attempt, err := DecryptAttempt(rawAttempt, []byte("password"))
	err = ioutil.WriteFile("../secureattempt", []byte(attempt.String()), 0600)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wrote decrypted to file")
}

func TestFromBytes(t *testing.T) {
	newAttempt, err := NewAttempt("username", 1234, []byte("payload"))
	if err != nil {
		t.Fatal(err)
	}

	constructedAttempt, err := AttemptFromBytes(newAttempt.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(constructedAttempt.String())

}
