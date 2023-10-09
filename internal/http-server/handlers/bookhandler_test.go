package handlers

import (
	"testing"

	".github.com/Luzik-D/BasicCRUD/internal/storage"
)

func TestValidatePOST(t *testing.T) {
	var test_cases = []struct {
		name  string
		input storage.Book
		want  error
	}{
		{"t1", storage.Book{Title: "x", Author: "y"}, nil},
		{"t2", storage.Book{Title: "x", Author: ""}, ErrInvalidPOSTRequest},
		{"t3", storage.Book{Title: "", Author: "y"}, ErrInvalidPOSTRequest},
		{"t4", storage.Book{Title: "", Author: ""}, ErrInvalidPOSTRequest},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			res := validatePOST(test.input)

			if res != test.want {
				t.Errorf("Test %s: got %v, expected %v", test.name, res, test.want)
			}
		})
	}
}

func TestValidatePUT(t *testing.T) {
	var test_cases = []struct {
		name  string
		input storage.Book
		want  error
	}{
		{"t1", storage.Book{Title: "x", Author: "y"}, nil},
		{"t2", storage.Book{Title: "x", Author: ""}, ErrInvalidPUTRequest},
		{"t3", storage.Book{Title: "", Author: "y"}, ErrInvalidPUTRequest},
		{"t4", storage.Book{Title: "", Author: "y"}, ErrInvalidPUTRequest},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			res := validatePUT(test.input)

			if res != test.want {
				t.Errorf("Test %s: got %v, expected %v", test.name, res, test.want)
			}
		})
	}
}

func TestValidatePATCH(t *testing.T) {
	var test_cases = []struct {
		name  string
		input storage.Book
		want  error
	}{
		{"t1", storage.Book{Title: "x", Author: "y"}, nil},
		{"t2", storage.Book{Title: "x", Author: ""}, nil},
		{"t3", storage.Book{Title: "", Author: "y"}, nil},
		{"t4", storage.Book{Title: "", Author: ""}, ErrInvalidPATCHRequest},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			res := validatePATCH(test.input)

			if res != test.want {
				t.Errorf("Test %s: got %v, expected %v", test.name, res, test.want)
			}
		})
	}
}
