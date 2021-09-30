package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	b := bytes.NewBuffer(nil)
	cmd := exec.Command("powershell.exe", "cat")
	cmd.Stdin = b
	cmd.Stdout = os.Stdout
	fmt.Fprintf(b, "Hello, World! I'm using this memory address: %p\n", b)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	cmd.Wait()
}