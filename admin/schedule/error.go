package schedule

import (
	"errors"
)


var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrCounselorNotFound = errors.New("counselor not found")
	ErrPageNotFound = errors.New("page not found")

	// bad request
	ErrIdFormat = errors.New("id required and must be a valid uuid")
)