package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <book-id>")
		return
	}
	catalog, err := books.OpenCatalog("testdata/catalog.json")
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
	}
	id := os.Args[1]
	book, ok := catalog.GetBook(id)
	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fmt.Println(book)
}
