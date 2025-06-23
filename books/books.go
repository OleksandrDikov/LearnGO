package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func GetCatalog() map[string]Book {
	return map[string]Book{
		"1": {
			ID:     "1",
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		"2": {
			ID:     "2",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
}

func BookToString(book Book) string {
	return fmt.Sprintf("%s by %s (copies: %d)", book.Title, book.Author, book.Copies)
}

func GetAllBooks(catalog map[string]Book) []Book {
	return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func AddBook(catalog map[string]Book, book Book) bool {
	catalog[book.ID] = book
	return true
}
