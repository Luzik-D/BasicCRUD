package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	".github.com/Luzik-D/BasicCRUD/internal/storage"
)

type BookHandler interface {
	//GetBookById(book int) (storage.Book, error)
	GetBooks() ([]storage.Book, error)
	GetBookById(id int) (storage.Book, error)
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func HandleBooks(storage BookHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := storage.GetBooks()
		if err != nil {
			panic(err)
		}

		fmt.Println(books)
	}
}

func HandleBook(storage BookHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/books/")

		if id == "" {
			fmt.Println("incorrect path")
			return
		}

		idVal, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("bad num")
			return
		}
		book, err := storage.GetBookById(idVal)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("FOUND: %v", book)
	}
}
