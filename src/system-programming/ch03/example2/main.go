package main

import (
  "fmt"
  "os"
  "os/exec"
)

func main() {
  cmd := exec.Command("ls", "-al")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  err := cmd.Run()
  if err != nil {
    fmt.Println(err)
    return
  }

  childPid := cmd.Process.Pid
  fmt.Println("Child process ID: ", childPid)

  pid := os.Getpid()
  fmt.Println("Current process ID: ", pid)
}
