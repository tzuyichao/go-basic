package main

import (
    "log"
    "runtime"
)

const info = `
  Application %s starting.
  The binary was build by Go: %s`

func main() {
    log.Printf(info, "Example", runtime.Version())
}