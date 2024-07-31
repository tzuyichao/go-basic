package main

import (
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    trimmed := strings.TrimSpace(s)
    words := strings.Fields(trimmed)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}

func main() {
    input := "  Hello   world this is Golang     "
    result := reverseWords(input)
    fmt.Println(result)
}
