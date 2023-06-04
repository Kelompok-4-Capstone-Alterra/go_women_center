package counselor

import (
	"errors"
)


var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// not found
	ErrCounselorNotFound = errors.New("counselor not found")
	ErrReviewNotFound = errors.New("review not found")
	ErrPageNotFound = errors.New("page not found")

	// conflict
	ErrCounselorConflict = errors.New("counselor already registered")
	ErrEmailConflict = errors.New("counselor email already registered")

	// bad request
	ErrProfilePictureFormat = errors.New("profile picture must be an png/jpg/jpeg and less than 2MB")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrTarifFormat = errors.New("tarif must be a number")
	ErrRatingFormat = errors.New("rating must be a number between 1-5")
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrInvalidTopic = errors.New("invalid topic")
	ErrRequired = errors.New("all fields are required")
)