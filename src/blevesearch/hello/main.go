package main

// https://blevesearch.com/docs/Getting%20Started/
import (
	"fmt"
	
    "github.com/blevesearch/bleve"
)

func main() {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	} {
		Name: "Text",
	}

	index.Index("id", data)

	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}