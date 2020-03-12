package main

import (
	"fmt"
	"time"
)

type Rocket struct {
    Name string
}

func (r *Rocket) Launch() {
    fmt.Printf("Rocket %s is launch....\n", r.Name)
}

func main() {
	// r := new(Rocket)
	r := Rocket {
		Name: "Test",
	}
	time.AfterFunc(10 * time.Second, func() {r.Launch()})
	fmt.Println("Completed.")
	time.Sleep(20 * time.Second)
}