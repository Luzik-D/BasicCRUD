package handlers

import "errors"

var (
	ErrInvalidPUTRequest   = errors.New("Invalid PUT request")
	ErrInvalidPOSTRequest  = errors.New("Invalid POST request")
	ErrInvalidPATCHRequest = errors.New("Invalid PATCH request")
)
