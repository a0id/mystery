package common

import "golang.org/x/crypto/sha3"

// DefaultPinLength - The default security pin length
var DefaultPinLength = 4

// Sha3 - Hash input using sha3
func Sha3(b []byte) []byte {
	hash := sha3.New256()
	hash.Write(b)
	return hash.Sum(nil)
}
