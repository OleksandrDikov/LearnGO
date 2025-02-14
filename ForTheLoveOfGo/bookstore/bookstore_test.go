package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	// First test
	want := bookstore.Book{ID: 1, Title: "For the Love of Go"}
	got := bookstore.GetBook(catalog, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	// Second test
	want = bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got = bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
