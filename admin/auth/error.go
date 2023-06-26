package auth

import "errors"

var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrInvalidCredential = errors.New("invalid credential")

	// conflict

	// bad request
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrRequired = errors.New("all fields are required")
)