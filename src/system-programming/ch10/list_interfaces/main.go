package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	for _, iface := range interfaces {
		fmt.Fprintf(os.Stdout, "Interface: %s\n", iface.Name)
	}
}
