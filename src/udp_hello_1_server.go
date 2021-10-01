package main

import (
	"os"
	"net"
	"log"
	"bytes"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address.")
	}
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("Listening:", os.Args[1], err)
	}
	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Println("<-", addr, "Message error:", err)
			continue
		}
		msg := bytes.TrimSpace(b[:n])
		log.Printf("<- %q from %s", msg, addr)
		for i, l := 0, len(msg); i<l/2; i++ {
			msg[i], msg[l-1-i] = msg[l-1-i], msg[i]
		}
		log.Printf("before write %q to %s\n", msg, addr)
		msg = append(msg, '\n')
		if _, err := conn.WriteTo(msg[:n], addr); err != nil {
			log.Println("->", addr, "Send error:", err)
		}
	}
}