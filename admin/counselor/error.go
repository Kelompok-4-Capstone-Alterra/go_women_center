package counselor

import (
	"errors"
)


var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrCounselorNotFound = errors.New("counselor not found")
	ErrPageNotFound = errors.New("page not found")

	// conflict
	ErrEmailConflict = errors.New("counselor email already registered")
	ErrUsernameConflict = errors.New("counselor username already registered")

	// bad request
	ErrProfilePictureFormat = errors.New("profile picture must be an png/jpg/jpeg and less than 2MB")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrPriceFormat = errors.New("price must be a number")
	ErrRatingFormat = errors.New("rating must be a number between 1-5")
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrInvalidTopic = errors.New("invalid topic")
	ErrRequiredTopic 		= errors.New("topic is required")
	ErrRequired = errors.New("all fields are required")
	ErrInvalidSort 			= errors.New("invalid sort")
	ErrHasScheduleFormat 	= errors.New("has_schedule must be true or false")
)