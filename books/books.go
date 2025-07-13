package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sync"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

type Catalog struct {
	mu   *sync.RWMutex
	data map[string]Book
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s (copies: %d)", b.Title, b.Author, b.Copies)
}

func (b *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	b.Copies = copies
	return nil
}

func NewCatalog() *Catalog {
	return &Catalog{
		mu:   &sync.RWMutex{},
		data: map[string]Book{},
	}
}

func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := NewCatalog()
	err = json.NewDecoder(file).Decode(&catalog.data)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

func (c *Catalog) GetAllBooks() []Book {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return slices.Collect(maps.Values(c.data))
}

func (c *Catalog) GetBook(id string) (Book, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	book, ok := c.data[id]
	return book, ok
}

func (c *Catalog) AddBook(book Book) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.data[book.ID]
	if ok {
		return fmt.Errorf("book with ID %s already exists", book.ID)
	}
	c.data[book.ID] = book
	return nil
}

func (c *Catalog) Sync(path string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(c.data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Catalog) SetCopies(ID string, copies int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	book, ok := c.data[ID]
	if !ok {
		return fmt.Errorf("ID %q not found", ID)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	c.data[ID] = book
	return nil
}

func (c *Catalog) GetCopies(ID string) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	book, ok := c.data[ID]
	if !ok {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	return book.Copies, nil
}
