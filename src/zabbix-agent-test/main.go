package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: zabbix_agent_test <host> <port> <item_key>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]
	itemKey := os.Args[3]
	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to Zabbix Agent: %v\n", err)
		return
	}
	defer conn.Close()

	header := []byte("ZBXD\x01")
	dataLength := uint64(len(itemKey))

	buffer := make([]byte, 13+len(itemKey))
	copy(buffer[0:5], header)
	binary.LittleEndian.PutUint64(buffer[5:13], dataLength)
	copy(buffer[13:], itemKey)

	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	response := make([]byte, 512)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Response from Zabbix Agent: %s\n", string(response[:n]))
}
