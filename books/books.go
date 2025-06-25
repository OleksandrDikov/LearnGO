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

type Catalog map[string]Book

func GetCatalog() Catalog {
	return Catalog{
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

func (b Book) String() string {
	return fmt.Sprintf("%s by %s (copies: %d)", b.Title, b.Author, b.Copies)
}

func (b *Book) SetCopies(copies int) error {
	if copies < 1 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	b.Copies = copies
	return nil
}

func (c Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(c))
}

func (c Catalog) GetBook(id string) (Book, bool) {
	book, ok := c[id]
	return book, ok
}

func (c Catalog) AddBook(book Book) bool {
	c[book.ID] = book
	return true
}
