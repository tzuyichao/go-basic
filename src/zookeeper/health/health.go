package main

import (
	"fmt"
	"log"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	zkAddresses := []string{"localhost:2181"}
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
