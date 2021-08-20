package main

import (
	"log"
	"net/http"
	"fmt"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (c *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.counter += 1
	heading := fmt.Sprintf("<h1>%s</h1>", c.heading)
	content := fmt.Sprintf("<p>%s</p>", c.content)
	counter := fmt.Sprintf("<p>Views: %d</p>", c.counter)
	res := fmt.Sprintf("%s\n%s\n%s", heading, content, counter)
	w.Write([]byte(res))
}

func main() {
	root := PageWithCounter{
		counter: 0,
		content: "This is main page",
		heading: "Hello World",
	}
	http.Handle("/", &root)
	chapter1 := PageWithCounter{
		counter: 0,
		content: "This is the first page",
		heading: "Chapter1",
	}
	http.Handle("/chapter1", &chapter1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}