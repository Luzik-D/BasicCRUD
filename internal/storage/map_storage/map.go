package map_storage

import (
	"fmt"

	".github.com/Luzik-D/BasicCRUD/internal/storage"
)

/*
Simple storage based on map.

	Key -- book's id
	Value -- book
*/

type Storage struct {
	st      map[int]storage.Book
	last_id int
}

func New() (*Storage, error) {
	st := make(map[int]storage.Book)

	return &Storage{st, 1}, nil
}

func (s *Storage) GetBooks() ([]storage.Book, error) {
	var books []storage.Book

	for _, b := range s.st {
		books = append(books, b)
	}

	return books, nil
}

func (s *Storage) AddBook(b storage.Book) error {
	b.Id = s.last_id
	s.st[s.last_id] = b
	s.last_id++

	return nil
}

func (s *Storage) DeleteBookById(id int) error {
	delete(s.st, id)

	return nil
}

func (s *Storage) GetBookById(id int) (storage.Book, error) {
	book, ok := s.st[id]
	if !ok {
		return storage.Book{}, fmt.Errorf("Book with id %d doesn't exist", id)
	}

	return book, nil
}

func (s *Storage) UpdateBookWithId(id int, changes storage.Book) error {
	book, ok := s.st[id]
	if !ok {
		return fmt.Errorf("Book with id %d doesn't exist", id)
	}

	book.Author = changes.Author
	book.Title = changes.Title

	s.st[id] = book

	return nil
}

func (s *Storage) PatchBookWithId(id int, changes storage.Book) error {
	book, ok := s.st[id]
	if !ok {
		return fmt.Errorf("Book with id %d doesn't exist", id)
	}

	if changes.Author != "" {
		book.Author = changes.Author
	}

	if changes.Title != "" {
		book.Title = changes.Title
	}

	s.st[id] = book

	return nil
}
