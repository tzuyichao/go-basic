package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <zk-address>\n", os.Args[0])
		os.Exit(1)
	}
	zkAddresses := []string{os.Args[1]}
	conn, _, err := zk.Connect(zkAddresses, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	state := conn.State()
	fmt.Printf("Cluster State: %s\n", state)

	folders, _, err := conn.Children("/")
	if err != nil {
		log.Fatal(err)
	}
	for _, folder := range folders {
		fmt.Printf("Folder: %s\n", folder)
	}
}
