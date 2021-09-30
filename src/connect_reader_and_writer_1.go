package main

import (
	"io"
	"fmt"
)

func main() {
	pr, pw := io.Pipe()
	go func(w io.WriteCloser) {
		for _, s := range []string{"a string", "another string", "last one"} {
			fmt.Printf("-> Writing %q\n", s)
			fmt.Fprint(w, s)
		}
		w.Close()
	}(pw)
	var err error
	for n, b := 0, make([]byte, 100); err == nil; {
		fmt.Println("<- Waiting...")
		n, err = pr.Read(b)
		if err == nil {
			fmt.Printf("<- Received %q\n", string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("error: ", err)
	}
}