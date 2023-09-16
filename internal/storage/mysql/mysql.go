package mysql

import (
	"database/sql"
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

func (s *Storage) GetAllBooks() ([]storage.Book, error) {
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
	return nil
}

func (s *Storage) GetBookByTitle(title string) (storage.Book, error) {
	return storage.Book{}, nil
}

func (s *Storage) GetBookByAuthor(author string) (storage.Book, error) {
	return storage.Book{}, nil
}
