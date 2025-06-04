package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	fmt.Println("Books in stock:")
	book := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	printBook(book)
}

func printBook(book Book) {
	fmt.Printf("%s by %s - %d copies\n", book.Title, book.Author, book.Copies)
}
