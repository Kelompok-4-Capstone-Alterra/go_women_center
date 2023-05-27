package counselor

import (
	"errors"
)


var (
	ErrInternalServerError = errors.New("internal server error")
	
	ErrNotFound = errors.New("counselor not found")

	ErrCounselorConflict = errors.New("counselor already registered")

	ErrEmailConflict = errors.New("counselor email already registered")

	// validation error
	ErrProfilePictureFormat = errors.New("profile picture must be an image and png/jpg/jpeg format")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrTarifFormat = errors.New("tarif must be a number")
	ErrIdFormat = errors.New("id must be a valid uuid")
	ErrInvalidTopic = errors.New("invalid topic")
	ErrRequired = errors.New("all fields are required")
)

