package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/a0id/mystery/common"
)

// ErrInvalidArgumentsNewAttempt - Error is thrown when NewAttempt params are blank
var ErrInvalidArgumentsNewAttempt = errors.New("arguments to call NewAttempt() must not be blank")

// Attempt - An attempt to login
type Attempt struct {
	Username  *Username `json:"username"`
	Payload   []byte    `json:"payload"`
	Origin    string    `json:"origin"`
	Timestamp string    `json:"timestamp"`
	Hash      []byte    `json:"hash"`
}

// NewAttempt - Create a new attempt struct
func NewAttempt(username string, pin int, payload []byte) (*Attempt, error) {
	if username == "" || payload == nil {
		return nil, ErrInvalidArgumentsNewAttempt
	}

	newUsername, err := NewUsername(username, pin)
	if err != nil {
		return nil, err
	}

	newOrigin, err := common.GetPublicIP()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	newAttempt := &Attempt{
		Username:  newUsername,
		Payload:   payload,
		Origin:    newOrigin,
		Timestamp: timestamp.String(),
	}

	(*newAttempt).Hash = common.Sha3(newAttempt.Bytes())
	return newAttempt, nil
}

// EncryptAttempt - Encrypt an attempt struct
func EncryptAttempt(rawAttempt Attempt, passphrase []byte) ([]byte, error) {
	attempt := rawAttempt.Bytes()
	secureAttempt, err := common.AESEncrypt(attempt, common.Sha3(passphrase))
	if err != nil {
		return nil, err
	}
	return secureAttempt, nil
}

// DecryptAttempt - Decrypt bytes and return an Attempt struct
func DecryptAttempt(encryptedAttempt []byte, passphrase []byte) (*Attempt, error) {
	// Decrypt the bytes
	rawAttempt, err := common.AESDecrypt(encryptedAttempt, common.Sha3(passphrase))
	if err != nil {
		return nil, err
	}

	// Construct an attempt
	attempt, err := AttemptFromBytes(rawAttempt)
	if err != nil {
		return nil, err
	}
	return attempt, nil
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

// AttemptFromBytes - Construct an Attempt struct from a []byte
func AttemptFromBytes(bytes []byte) (*Attempt, error) {
	buffer := &Attempt{}
	err := json.Unmarshal(bytes, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
