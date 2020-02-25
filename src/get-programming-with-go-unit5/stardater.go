package main

import (
    "fmt"
    "time"
)

type stardater interface {
    YearDay() int
    Hour() int
}

func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour())/24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	s := sol(1422)
	fmt.Println(stardate(s))

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed.\n", stardate(day))
}