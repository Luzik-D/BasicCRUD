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
		id INT PRIMARY KEY,
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
