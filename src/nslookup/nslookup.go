package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <domain>\n", os.Args[0])
		os.Exit(1)
	} else {
		target := os.Args[1]
		ips, err := net.LookupIP(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
			os.Exit(1)
		}
		for _, ip := range ips {
			fmt.Printf("%s In A %s\n", target, ip)
		}
	}
}
