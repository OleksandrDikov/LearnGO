package main

import (
	"books"
	"fmt"
)

func main() {
	fmt.Println("Books in stock:")
	book := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	fmt.Println(books.BookToString(book))
}
