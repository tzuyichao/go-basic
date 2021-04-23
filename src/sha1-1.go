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

func main() {
    var bytes []byte
    bytes = CreateSHA1Hash([]byte("Check"))
    result := fmt.Sprintf("%x", bytes)
    fmt.Println(result)
}