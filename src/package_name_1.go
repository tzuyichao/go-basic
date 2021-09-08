package main

import (
	crand "crypto/rand"
	"math/rand"
	"fmt"
	"encoding/binary"
)

// from https://play.golang.org/p/QD8bqT38Zot
func main() {
	r := seedRand()
	fmt.Println(r.Int())
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}