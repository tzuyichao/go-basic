package main

import (
	"fmt"
	"io/ioutil"
	"log"

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
	knownHostsPath := "/path/to/your/known_hosts"
	targetHost := "your_target_host_name_or_ip"

	key, err := findPublicKey(knownHostsPath, targetHost)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Public key for %s: %s\n", targetHost, ssh.FingerprintSHA256(*key))
}
