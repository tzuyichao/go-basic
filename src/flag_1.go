package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ArrayValue) Set(s string) error {
	*s = strings.Split(s, ",")
	return nil
}

func main() {
	
}