package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/a0id/mystery/common"
)

// ErrInvalidArgumentsNewAttempt - Error is thrown when NewAttempt params are blank
var ErrInvalidArgumentsNewAttempt = errors.New("arguments to call NewAttempt must not be blank")

// Attempt - An attempt to login
type Attempt struct {
	Username  *Username `json:"username"`
	Token     []byte    `json:"token"`
	Origin    string    `json:"origin"`
	Timestamp string    `json:"timestamp"`
	Hash      []byte    `json:"hash"`
}

// NewAttempt - Create a new attempt struct
func NewAttempt(username string, pin int, token string) (*Attempt, error) {
	if username == "" || token == "" {
		return nil, ErrInvalidArgumentsNewAttempt
	}

	newUsername, err := NewUsername(username, pin)
	if err != nil {
		return nil, err
	}

	newToken := []byte(token)

	newOrigin, err := common.GetPublicIP()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	newAttempt := &Attempt{
		Username:  newUsername,
		Token:     newToken,
		Origin:    newOrigin,
		Timestamp: timestamp.String(),
	}

	(*newAttempt).Hash = common.Sha3(newAttempt.Bytes())

}

// Bytes - Encode an attempt into a []byte
func (attempt *Attempt) Bytes() []byte {
	json, _ := json.MarshalIndent(*attempt, "", "  ")
	return json
}

// String - Encode an attempt into a string
func (attempt *Attempt) String() string {
	json, _ := json.MarshalIndent(*attempt, "", "  ")
	return string(json)
}
