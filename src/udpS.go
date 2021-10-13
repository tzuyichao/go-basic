package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	port := "127.0.0.1:" + arguments[1]
	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	connect, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connect.Close()
	fmt.Println("Server Ready.")
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())
	for {
		n, addr, err := connect.ReadFromUDP(buffer)
		fmt.Print(addr.IP.String(), ":", addr.Port, " -> ", string(buffer[0:n-1]))
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP Server")
			return
		}
		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("%s:%d data: %s\n", addr.IP.String(), addr.Port, string(data))
		_, err = connect.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}