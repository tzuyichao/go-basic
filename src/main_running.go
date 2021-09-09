package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "1")
	proc.Start()

	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)
}