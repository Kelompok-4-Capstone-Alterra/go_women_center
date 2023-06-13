package users

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")

	ErrUserNotFound = errors.New("user not found")

	ErrIdFormat = errors.New("id must be a valid uuid")

	ErrIdRequired = errors.New("id is required")

	ErrPageNotFound = errors.New("page not found")
)