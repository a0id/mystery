package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// GenerateRSAKeypair - Generate and export an RSA public/private keypair
func GenerateRSAKeypair() error {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return err
	}

	publicKey := key.PublicKey

	// Export the keys
	err = ExportPEMKey("private.pem", key)
	if err != nil {
		return err
	}
	err = ExportPEMPublicKey("public.pem", publicKey)
	if err != nil {
		return err
	}
	return nil
}

// ExportPEMKey - Write a private key to memory in PEM format
func ExportPEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}
	return nil
}

// ExportPEMPublicKey - Write a public key to memory in PEM format
func ExportPEMPublicKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}
	return nil
}

// EncryptRSA - Encrypt a []byte with RSA
func EncryptRSA(input []byte, outFileName string) error {
	// Read the private key
	pemData, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return err
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return errors.New("bad key data: not PEM-encoded")
	}

	if block.Type != "PRIVATE KEY" {
		return errors.New("pem block is not an rsa private key")
	}

	// Decode the RSA private key
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	// Set up the encryption stuff
	randHash := sha1.New()
	randReader := rand.Reader
	// randReader := rand.Reader
	fmt.Printf("randHash: %x\n", &randHash)

	// Encrypt the data
	encrypted, err := rsa.EncryptOAEP(randHash, randReader, &priv.PublicKey, input, []byte(outFileName))
	if err != nil {
		return err
	}

	// Write data to output file
	err = ioutil.WriteFile(outFileName, []byte(encrypted), 0600)
	if err != nil {
		return err
	}

	// Decrypt the data
	decrypted, err := rsa.DecryptOAEP(randHash, nil, priv, input, []byte(outFileName))
	if err != nil {
		return err
	}
	fmt.Println("decrypted: " + string(decrypted))

	// Write data to output file
	// err = ioutil.WriteFile(outFileName, []byte(decrypted), 0600)
	// if err != nil {
	// 	return err
	// }
	return nil
}
