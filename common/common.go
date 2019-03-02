package common

import (
	"net"
	"strings"

	"github.com/glendc/go-external-ip"
	"golang.org/x/crypto/sha3"
)

// DefaultPinLength - The default security pin length
var DefaultPinLength = 4

// BuffSize - The server's buffer size
var BuffSize = 1024

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
