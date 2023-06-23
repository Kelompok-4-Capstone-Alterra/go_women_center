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
	ErrScheduleNotFound = errors.New("counselor doesn't have schedule today")
	ErrPageNotFound = errors.New("page not found")

	// bad request
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrDateInvalid = errors.New("date must be a valid date")
	ErrTimeInvalid = errors.New("time must be a valid time")
)