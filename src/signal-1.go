package main

//go build -o "signal-1" signal-1.go
import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("Start application...")
	c := make(chan os.Signal)
	signal.Notify(c)
	s := <-c
	log.Println("Exit with signal:", s)
}
