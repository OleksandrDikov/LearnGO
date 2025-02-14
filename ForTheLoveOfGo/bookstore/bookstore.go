package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, id int) Book {
	for _, v := range catalog {
		if v.ID == id {
			return v
		}
	}
	return Book{}
}
