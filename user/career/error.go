package career

import (
	"errors"
)

var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrCareerNotFound = errors.New("career not found")
	ErrPageNotFound   = errors.New("page not found")

	// bad request
	ErrIdFormat    = errors.New("id must be a valid uuid")
	ErrRequired    = errors.New("all fields are required")
	ErrInvalidSort = errors.New("invalid sort")
)
