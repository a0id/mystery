package types

import (
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

	secureAttempt, err := EncryptAttempt(attempt, []byte("password"))
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("../secureattempt", secureAttempt, 0600)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wrote to file")
}

func TestDecrypt(t *testing.T) {
	rawAttempt, err := ioutil.ReadFile("../secureattempt")
	if err != nil {
		t.Fatal(err)
	}

	attempt, err := DecryptAttempt(rawAttempt, []byte("password"))
	t.Logf("decrypted: %s\n", attempt.String())
}
