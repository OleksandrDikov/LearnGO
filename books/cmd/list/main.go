package main

import (
	"books"
	"fmt"
)

func main() {
	catalog, err := books.OpenCatalog("testdata/catalog.json")
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
	}
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book)
	}
}
