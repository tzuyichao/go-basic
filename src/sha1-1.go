package main

import (
    "fmt"
    "crypto/sha1"
    "hash"
)

func CreateSHA1Hash(byteStr []byte) []byte {
    var hashVal hash.Hash
    hashVal = sha1.New()
    hashVal.Write(byteStr)

    var bytes []byte
    bytes = hashVal.Sum(nil)
    return bytes
}

func ExecSHA1Hash(txt string) {
    var bytes []byte
    bytes = CreateSHA1Hash([]byte(txt))
    result := fmt.Sprintf("%x", bytes)
    fmt.Printf("%s: %s\n", txt, result)
}

func main() {
	ExecSHA1Hash("Check")
	ExecSHA1Hash("check")
}
