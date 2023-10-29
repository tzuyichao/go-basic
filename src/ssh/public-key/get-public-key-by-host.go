package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func findPublicKey(knownHostsPath, targetHost string) (*ssh.PublicKey, error) {
	publicKeyData, err := ioutil.ReadFile(knownHostsPath)
	if err != nil {
		return nil, err
	}

	_, hosts, key, _, _, err := ssh.ParseKnownHosts(publicKeyData)
	if err != nil {
		return nil, err
	}

	for _, host := range hosts {
		if host == targetHost {
			return &key, nil
		}
	}

	return nil, fmt.Errorf("no public key found for %s", targetHost)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <path_to_known_hosts> <target_host>", os.Args[0])
	}

	knownHostsPath := os.Args[1]
	targetHost := os.Args[2]

	key, err := findPublicKey(knownHostsPath, targetHost)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Public key for %s: %s\n", targetHost, ssh.FingerprintSHA256(*key))
}
