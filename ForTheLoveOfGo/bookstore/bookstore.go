package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, v := range c {
		result = append(result, v)
	}
	return result
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	// return int(float64(b.PriceCents) * (1 - float64(b.DiscountPercent) / 100))
	return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
}