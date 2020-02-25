package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func (p *person) birthday2() {
    p.age++
}

func main() {
    terry := &person {
        name: "Terry",
        age: 15,
	}
	terry2 := terry
	terry.birthday2()
	fmt.Println(*terry)
	fmt.Println(*terry2)

	eli := person {
		name: "Eli",
		age: 20,
	}
	eli2 := eli
	eli2.birthday2()
	fmt.Println(eli)
	fmt.Println(eli2)
}