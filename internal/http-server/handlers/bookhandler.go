package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	".github.com/Luzik-D/BasicCRUD/internal/storage"
)

type BookHandler interface {
	GetBooks() ([]storage.Book, error)
	GetBookById(id int) (storage.Book, error)
	UpdateBookWithId(id int, changes storage.Book) error
	PatchBookWithId(id int, changes storage.Book) error
	AddBook(b storage.Book) error
	DeleteBookById(id int) error
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func validatePOST(b storage.Book) error {
	if b.Title == "" || b.Author == "" {
		return ErrInvalidPOSTRequest
	}

	return nil
}

func validatePUT(b storage.Book) error {
	if b.Title == "" || b.Author == "" {
		return ErrInvalidPUTRequest
	}

	return nil
}

func validatePATCH(b storage.Book) error {
	if b.Title == "" && b.Author == "" {
		return ErrInvalidPATCHRequest
	}

	return nil
}

func HandleBooks(stor BookHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			books, err := stor.GetBooks()
			if err != nil {
				log.Println("Failed to get books")
			}

			log.Println(books)
		case "POST":
			var book storage.Book
			// decode json
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Println("Failed to decode request body")
				return
			}

			// validate req
			err = validatePOST(book)
			if err != nil {
				log.Println("Invalid POST request")
				return
			}

			// add to storage
			err = stor.AddBook(book)
			if err != nil {
				log.Println("Failed to add the book")
				return
			}

			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandleBook(stor BookHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/books/")

		if id == "" {
			log.Println("Empty id")
			return
		}

		idVal, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Failed to convert id from path")
			return
		}
		book, err := stor.GetBookById(idVal)
		if err != nil {
			log.Println("Failed to get book from storage")
			return
		}

		switch r.Method {
		case "GET":
			js, err := json.Marshal(book)
			if err != nil {
				log.Println("Failed to marshall the book")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case "PUT":
			var b storage.Book
			err := json.NewDecoder(r.Body).Decode(&b)
			if err != nil {
				log.Println("Failed to decode request body")
				return
			}

			err = validatePUT(b)
			if err != nil {
				log.Println("Invalid PUT request")
				return
			}

			qerr := stor.UpdateBookWithId(idVal, b)
			if qerr != nil {
				log.Println("Failed to update the book")
				return
			}
			w.Write([]byte("ok"))
		case "PATCH":
			var b storage.Book
			err := json.NewDecoder(r.Body).Decode(&b)
			if err != nil {
				log.Println("Failed to decode request body")
				return
			}

			err = validatePATCH(b)
			if err != nil {
				log.Println("Invalid PATCH request")
				return
			}

			qerr := stor.PatchBookWithId(idVal, b)
			if qerr != nil {
				log.Println("Failed to patch the book")
				return
			}
			w.Write([]byte("ok"))
		case "DELETE":
			err := stor.DeleteBookById(idVal)
			if err != nil {
				log.Println("Failed to delete the book")
			}
			w.Write([]byte("ok"))
		default:
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}
