package main

import (
	"context"
	"fmt"
)

func logic(ctx context.Context, info string) (string, error) {
	fmt.Println("context:", ctx)
	fmt.Println("info:", info)
	return "", nil
}

func main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", result)
}