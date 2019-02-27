package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

// DefaultPinLength - The default security pin length
var DefaultPinLength = 4

// ErrBadKeyData - Error is returned when a key is not PEM-encoded
var ErrBadKeyData = errors.New("bad key data: not PEM-encoded")

// ErrOtherKeyErr - Unspecified / Random / other key error
var ErrOtherKeyErr = errors.New("other key error")

// Sha3 - Hash input using sha3
func Sha3(b []byte) []byte {
	hash := sha3.New256()
	hash.Write(b)
	return hash.Sum(nil)
}

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

// GetPublicIP - Get the host's public IP
func GetPublicIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx], nil
}

// EncryptRSA - Encrypt a []byte with RSA
func EncryptRSA(in []byte) error {
	// in := []byte("hello this is test")

	// Read the private key
	pemData, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return err
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return ErrBadKeyData
	}

	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		return ErrOtherKeyErr
	}

	// Decode the RSA private key
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return ErrOtherKeyErr
	}

	outFileName := "test.txt"

	var out []byte
	// Decrypt the data
	out, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, in, []byte(outFileName))
	if err != nil {
		log.Fatalf("decrypt: %s", err)
	} else {
		out, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, &priv.PublicKey, in, []byte(outFileName))
		if err != nil {
			log.Fatalf("encrypt: %s", err)
		}
	}

	// Write data to output file
	if err := ioutil.WriteFile(*outFile, out, 0600); err != nil {
		log.Fatalf("write output: %s", err)
	}

}
