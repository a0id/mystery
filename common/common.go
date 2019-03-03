package common

import (
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	externalip "github.com/glendc/go-external-ip"
	"golang.org/x/crypto/sha3"
)

// DefaultPinLength - The default security pin length
var DefaultPinLength = 4

// BuffSize - The server's buffer size
var BuffSize = 1000000 // 1 MB

// ExportDir - The dir to export accepted payloads
var ExportDir = "data/"

// NEWLINE - A universal newline character
var NEWLINE = "\n"

// InitCommon - Initialize the common package
func InitCommon() {
	if runtime.GOOS == "windows" {
		NEWLINE = "\r\n" // The newline character on Windows
	}
}

// Sha3 - Hash input using sha3
func Sha3(b []byte) []byte {
	hash := sha3.New256()
	hash.Write(b)
	return hash.Sum(nil)
}

// GetPublicIP - Get the host's public IP
func GetPublicIP() (string, error) {
	// Get the local IP address
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	// The final local ip
	localIP := localAddr[0:idx]

	// Get the public IP
	consensus := externalip.DefaultConsensus(nil, nil)

	// Get the host's ip
	ip, err := consensus.ExternalIP()
	if err != nil {
		return "", err
	}

	return ip.String() + "/" + localIP, nil
}

// CreateDirIfDoesNotExit - create given directory if does not exist
func CreateDirIfDoesNotExit(dir string) error {
	dir = filepath.FromSlash(dir) // Just to be safe

	if _, err := os.Stat(dir); os.IsNotExist(err) { // Check dir exists
		err = os.MkdirAll(dir, 0755) // Create directory

		if err != nil { // Check for errors
			return err // Return error
		}
	}

	return nil // No error occurred
}
