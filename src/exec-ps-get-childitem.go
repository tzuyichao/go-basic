package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	process := exec.Command("powershell", "Get-ChildItem", "-Name")
	output := bytes.NewBuffer([]byte{})
	process.Stdout = output
	err := process.Run()
	if err != nil {
		fmt.Println(err)
	}
	if process.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(output.String())
	}
}