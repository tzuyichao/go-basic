package main

import (
    "bytes"
    "testing"
)

func TestCountWords(t *testing.T) {
    b := bytes.NewBufferString("word1 word2 word3 word4\n")
    expect :=4
    res := count(b)
    if res != expect {
        t.Errorf("Expected %d, got %d instead.\n", expect, res)
    }
}