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

func (s *ArrayValue) Set(str string) error {
	*s = strings.Split(str, ",")
	return nil
}

func main() {
	retry := flag.Int("retry", -1, "Defines max retry count")
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger Prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through")

	flag.Parse()

	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying Connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}