package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
