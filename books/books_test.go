package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
		"1": {
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "1",
		},
		"2": {
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "2",
		},
	}
}

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []books.Book{
		{
			ID:     "1",
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			ID:     "2",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}

	got := books.GetAllBooks(catalog)
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book{
		ID:     "1",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := books.GetBook(catalog, "1")
	if !ok {
		t.Fatalf("Book not found.")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book{
		ID:     "3",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	}
	_, ok := books.GetBook(catalog, want.ID)
	if ok {
		t.Fatalf("Book already exists.")
	}

	ok = books.AddBook(catalog, want)
	if !ok {
		t.Fatalf("Error adding book.")
	}
}
