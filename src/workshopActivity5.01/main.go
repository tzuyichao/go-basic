package main

import (
	"fmt"
)

type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHour(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for i := 0; i<7; i++ {
		total = total + d.WorkWeek[i]
	}
	return total
}

func main() {
	var d1 Developer
	d1.LogHour(Monday, 8)
	d1.LogHour(Tuesday, 10)
	fmt.Printf("Hours worked on Monday: %d\n", d1.WorkWeek[Monday])
	fmt.Printf("Hours worked on Tuesday: %d\n", d1.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week: %d\n", d1.HoursWorked())
}