package main

import (
	"io"
	"io/ioutil"
	"unicode/utf8"
	"unicode"
	"fmt"
	"log"
	"strings"
)

func NewAngryReader(r io.Reader) *AngryReader {
	return &AngryReader{r: r}
}

type AngryReader struct {
	r io.Reader
}

func (a *AngryReader) Read(b []byte) (int, error) {
	n, err := a.r.Read(b)
	for r, i, w := rune(0), 0, 0; i<n; i+=w {
		r, w = utf8.DecodeRune(b[i:])
		if !unicode.IsLetter(r) {
			continue
		}
		ru := unicode.ToUpper(r)
		if wu := utf8.EncodeRune(b[i:], ru); w != wu {
			return n, fmt.Errorf("%c->%c, size mismatch %d->%d", r, ru, w, wu)
		}
	}
	return n, err
}

func main() {
	a := NewAngryReader(strings.NewReader("Hello, playground!"))
	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
}