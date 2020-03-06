package main

import (
    "fmt"
	"os"
	"net/http"
	"golang.org/x/net/html"
	"strings"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [url]\n", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Printf("Getting title from %s\n", url)
	err := title(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	contentType := resp.Header.Get("Content-Type")
	log.Printf("Content-Type: %s\n", contentType)
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, contentType)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parse %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}