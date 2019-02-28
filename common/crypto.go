package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

func NewEncrypt(data []byte, passphrase string) ([]byte, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher([]byte(Sha3([]byte(passphrase))))
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

func NewDecrypt(data []byte, passphrase string) ([]byte, error) {
	// Create the cipher from passphrase
	key := []byte(Sha3([]byte(passphrase)))
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

func EncryptFile(filename string, data []byte, passphrase string) error {
	f, _ := os.Create(filename)
	defer f.Close()
	encrypted, err := NewEncrypt(data, passphrase)
	if err != nil {
		return err
	}
	f.Write(encrypted)
	return nil
}

func DecryptFile(filename string, passphrase string) ([]byte, error) {
	data, _ := ioutil.ReadFile(filename)
	return NewDecrypt(data, passphrase)
}
