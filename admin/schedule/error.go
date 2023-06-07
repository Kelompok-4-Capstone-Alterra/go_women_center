package schedule

import (
	"errors"
)


var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// required
	ErrIdRequired = errors.New("id required")
	ErrDatesRequired = errors.New("dates required")
	ErrTimesRequired = errors.New("times required")

	// not found
	ErrCounselorNotFound = errors.New("counselor not found")
	ErrPageNotFound = errors.New("page not found")

	// bad request
	ErrIdFormat = errors.New("id must be a valid uuid")
)