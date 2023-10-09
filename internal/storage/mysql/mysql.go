package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	".github.com/Luzik-D/BasicCRUD/internal/storage"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

// todo: add secure db opening
// todo: add indexes
// todo: add connection pool
// todo: log errors
func New() (*Storage, error) {
	db, err := sql.Open("mysql", "sunrise:141018@/BasicCRUD")
	if err != nil {
		return nil, fmt.Errorf("Failed to open DB: %s", err)
	}

	q, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS Book (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		title VARCHAR(50) NOT NULL,
		author VARCHAR(30) NOT NULL
	);`)

	if err != nil {
		return nil, fmt.Errorf("mysql New: %s", err)
	}

	_, qerr := q.Exec()
	if err != nil {
		return nil, fmt.Errorf("mysql New: %s", qerr)
	}

	return &Storage{db}, nil
}

func (s *Storage) GetBooks() ([]storage.Book, error) {
	q, err := s.db.Query("SELECT * FROM Book")

	if err != nil {
		return nil, fmt.Errorf("mysql GetAllBooks: %s", err)
	}

	var books []storage.Book

	for q.Next() {
		var book storage.Book
		err := q.Scan(&book.Id, &book.Title, &book.Author)

		if err != nil {
			return nil, fmt.Errorf("mysql GetAllBooks: %s", err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *Storage) AddBook(b storage.Book) error {
	q, err := s.db.Prepare("INSERT INTO Book (title, author) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("mysql AddBook: %s", err)
	}

	_, qerr := q.Exec(b.Title, b.Author)
	if qerr != nil {
		return fmt.Errorf("mysql AddBook: %s", qerr)
	}

	return nil
}

func (s *Storage) DeleteBookById(id int) error {
	q, err := s.db.Prepare("DELETE FROM Book WHERE id = ?")
	if err != nil {
		return fmt.Errorf("mysql DeleteBookById: %s", err)
	}

	_, qerr := q.Exec(id)
	if qerr != nil {
		return fmt.Errorf("mysql DeleteBookById: %s", qerr)
	}

	return nil
}

func (s *Storage) GetBookById(id int) (storage.Book, error) {
	q, err := s.db.Query("SELECT * FROM Book WHERE id = ?", id)
	if err != nil {
		return storage.Book{}, fmt.Errorf("mysql GetBookById: %s", err)
	}

	fmt.Println("id ", id)
	var book storage.Book

	var qerr error
	if q.Next() {
		qerr = q.Scan(&book.Id, &book.Title, &book.Author)
	} else {
		return storage.Book{}, fmt.Errorf("mysql GetBookById: %s", qerr)
	}

	if qerr != nil {
		return storage.Book{}, fmt.Errorf("mysql GetBookById: %s", qerr)
	}

	return book, nil
}

func (s *Storage) UpdateBookWithId(id int, changes storage.Book) error {
	fmt.Printf("CHANGES: %d, %s, %s\n", id, changes.Title, changes.Author)
	_, err := s.db.Query("UPDATE Book SET title = ?, author = ? WHERE id = ?", changes.Title, changes.Author, id)
	if err != nil {
		return fmt.Errorf("mysql UpdateBookWithId error: %s", err)
	}

	return nil
}

func (s *Storage) PatchBookWithId(id int, changes storage.Book) error {
	return errors.New("Not implemented")
}
