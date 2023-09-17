package handlers

import (
	"encoding/json"
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
	UpdateBookWithId(id int, changes storage.Book) error
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

func HandleBook(stor BookHandler) http.HandlerFunc {
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
		book, err := stor.GetBookById(idVal)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("FOUND: %v", book)

		switch r.Method {
		case "GET":
			js, err := json.Marshal(book)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case "PUT":
			var b storage.Book
			err := json.NewDecoder(r.Body).Decode(&b)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Parse body from PUT req ", b)
			qerr := stor.UpdateBookWithId(idVal, b)
			if qerr != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte("ok"))
		default:
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}
