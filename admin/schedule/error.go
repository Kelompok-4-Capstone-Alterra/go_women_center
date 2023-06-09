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
	ErrScheduleNotFound = errors.New("counselor doesn't have schedule")

	// bad request
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrDateInvalid = errors.New("date must be a valid date")
	ErrTimeInvalid = errors.New("time must be a valid time")
	ErrDatesOrTimesRequired = errors.New("dates or times required")
	ErrScheduleAlreadyExist = errors.New("schedule already exist")
)