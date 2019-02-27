package types

import (
	"encoding/json"
	"errors"

	"github.com/a0id/mystery/common"
)

// ErrInvalidUsername - The error thrown when the username is blank
var ErrInvalidUsername = errors.New("username cannot be blank")

// Username - A username
type Username struct {
	Username string `json:"username"`
	Pin      *Pin   `json:"pin"`
	Hash     []byte `json:"hash"`
}

// NewUsername - Create a new username
func NewUsername(username string, pin int) (*Username, error) {
	if username != "" {
		newPin, err := NewPin(pin)
		if err != nil {
			return nil, err
		}
		newUsername := &Username{
			Username: username,
			Pin:      newPin,
		}
		newUsername.Hash = common.Sha3(newUsername.Bytes())
		return newUsername, nil
	}
	return nil, ErrInvalidUsername
}

// Bytes - Encode a username into a []byte
func (username *Username) Bytes() []byte {
	json, _ := json.MarshalIndent(*username, "", "  ")
	return json
}

// String - Encode a username into a string
func (username *Username) String() string {
	json, _ := json.MarshalIndent(*username, "", "  ")
	return string(json)
}
