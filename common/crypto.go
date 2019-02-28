package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

// AESEncrypt - Encrypt using an AES cipher
func AESEncrypt(data []byte, passphrase []byte) ([]byte, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher([]byte(Sha3(passphrase)))
	if err != nil {
		return nil, err
	}

	// Create the nonce
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	// Write the nonce to memory
	err = ioutil.WriteFile("nonce.txt", nonce, 0600)
	if err != nil {
		return nil, err
	}

	// Return the encrypted data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// AESDecrypt - Decrypt using an AES cipher
func AESDecrypt(data []byte, passphrase []byte) ([]byte, error) {
	// Create the cipher from passphrase
	key := []byte(Sha3(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Load the nonce
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()

	// Decrypt the data
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
