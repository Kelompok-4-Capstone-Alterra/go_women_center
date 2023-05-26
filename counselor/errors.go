package counselor

import (
	"errors"
)


var (
	ErrInternalServerError = errors.New("internal server error")
	
	ErrNotFound = errors.New("counselor not found")

	ErrConflict = errors.New("counselor already registered")

	

	// validation error
	ErrProfilePictureFormat = errors.New("profile picture must be an image and png/jpg/jpeg format")
	ErrEmailFormat = errors.New("email must be a valid email")
	ErrTarifFormat = errors.New("tarif must be a number")
	ErrRequired = errors.New("all fields are required")
	ErrInvalidTopic = errors.New("invalid topic")
)

