package main

import (
    "fmt"
)

func main() {
    amount, err := paintNeeded(4.2, -3.0)
    if err != nil {
        fmt.Println(err)
		return
    }
    fmt.Printf("%0.2f leters needed\n", amount)
}

func paintNeeded(width float64, height float64) (float64, error) {
    if width < 0 {
        return 0, fmt.Errorf("a width of %0.2f is invalud", width)
    }
    if height < 0 {
        return 0, fmt.Errorf("a height of %0.2f is invalud", height)
    }
    area := width * height
    return area / 10.0, nil
}