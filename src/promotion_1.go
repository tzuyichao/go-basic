package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	fmt.Println("Called")
	return []Employee{}
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID: "1234",
		},
		Reports: []Employee{},
	}
	fmt.Println(m)
	fmt.Println(m.Description())
}