package main

import (
	"books"
	"fmt"
)

func main() {
	catalog := books.GetCatalog()
	fmt.Println("Books in stock:")
	booksList := books.GetAllBooks(catalog)
	for _, book := range booksList {
		fmt.Println(books.BookToString(book))
	}
}
