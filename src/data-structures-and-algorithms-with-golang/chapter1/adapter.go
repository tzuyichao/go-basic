package main

import (
	"fmt"
)

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method: ", adaptee.adapterType)
}

func main() {
	var processor IProcess = Adapter{}

	processor.process()
}