package main

import (
    "fmt"
    "os"
)

func main() {
	args := os.Args
	
	fmt.Println(args)

	programName := args[0]
	fmt.Printf("The binary name is: %s\n", programName)

	otherArgs := args[1:]
	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}
}