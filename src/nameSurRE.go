package main 

import (
	"fmt"
	"regexp"
	"os"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: nameSurRE [text]")
		return
	}
	fmt.Println(matchNameSur(os.Args[1]))
}