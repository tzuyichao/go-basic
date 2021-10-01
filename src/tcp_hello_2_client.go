package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specifiy an address.")
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}
	createConn(addr)
}

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln("-> Connection:", err)
	}
	log.Println("-> Connection to", addr)
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("# ")
		msg, err := r.ReadBytes('\n')
		if err != nil {
			log.Println("-> Message error:", err)
		}
		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connectoion:", err)
			return
		}
	}
}