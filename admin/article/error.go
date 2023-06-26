package article

import (
	"errors"
)

var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrArticleNotFound = errors.New("article not found")
	ErrPageNotFound    = errors.New("page not found")
	ErrCommentNotFound = errors.New("comment not found")

	// bad request
	ErrImageFormat = errors.New("image must be an png/jpg/jpeg and less than 2MB")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrIdFormat    = errors.New("id must be a valid uuid")
	ErrRequired    = errors.New("all fields are required")
	ErrInvalidSort = errors.New("invalid sort")
)
