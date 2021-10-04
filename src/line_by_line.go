package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println("End Of File")
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Println(strings.TrimSpace(line))
	}
	return nil
}

func main() {
	lineByLine("./csv.data")
}