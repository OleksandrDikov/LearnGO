package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) map[int]Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	if _, ok := catalog[id]; !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return catalog[id], nil
}
