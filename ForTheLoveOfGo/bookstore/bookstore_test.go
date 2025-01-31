package bookstore_test

import (
	"bookstore"
	"testing"
)

//
//bookstore.Book {
//	Title: "The Art of Computer Programming",
//	Author: "Donald Knuth",
//	Copies: 3,
//}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}
