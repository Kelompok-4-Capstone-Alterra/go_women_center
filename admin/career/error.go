package career

import (
	"errors"
)

var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrCareerNotFound = errors.New("career not found")

	// bad request
	ErrImageFormat = errors.New("image must be an image and png/jpg/jpeg format")
	ErrImageSize = errors.New("image size must be less than 10MB")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrRequired = errors.New("all fields are required")
)