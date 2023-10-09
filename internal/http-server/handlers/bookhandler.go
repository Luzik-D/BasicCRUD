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

// Handler on "http://localhost:{PORT}/books"
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

// Handle on "http://localhost:{PORT}/books/{id}"
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

		switch r.Method {
		case "GET":
			getBookHandler(idVal, stor, w)
		case "PUT":
			putBookHandler(idVal, stor, r)
		case "PATCH":
			patchBookHandler(idVal, stor, r)
		case "DELETE":
			deleteBookHandler(idVal, stor)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

func getBookHandler(id int, st BookHandler, w http.ResponseWriter) error {
	book, err := st.GetBookById(id)
	if err != nil {
		log.Println("Failed to get book from storage")
		return err
	}

	js, err := json.Marshal(book)
	if err != nil {
		log.Println("Failed to marshall book")
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func putBookHandler(id int, st BookHandler, r *http.Request) error {
	var b storage.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("Failed to decode request")
		return err
	}

	err = validatePUT(b)
	if err != nil {
		log.Println("Invalid PUT request")
	}

	err = st.UpdateBookWithId(id, b)
	if err != nil {
		log.Println("Failed to update book")
		return err
	}

	return nil
}

func patchBookHandler(id int, st BookHandler, r *http.Request) error {
	var b storage.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("Failed to decode request")
		return err
	}

	err = validatePATCH(b)
	if err != nil {
		log.Println("Invalid PATCH request")
	}

	err = st.PatchBookWithId(id, b)
	if err != nil {
		log.Println("Failed to patch book")
		return err
	}

	return nil
}

func deleteBookHandler(id int, st BookHandler) error {
	err := st.DeleteBookById(id)
	if err != nil {
		log.Println("Failed to delete book")
	}

	return nil
}
