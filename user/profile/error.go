package profile

import "errors"

var (
	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	// duplicate
	ErrEmailDuplicate = errors.New("email already registered")
	ErrUsernameDuplicate = errors.New("username already registered")

	// not found
	ErrUserNotFound = errors.New("user not found")
	ErrPageNotFound = errors.New("page not found")

	// bad request
	ErrProfilePictureFormat = errors.New("profile picture must be an image and png/jpg/jpeg format")
	ErrEmailFormat          = errors.New("email must be a valid email")
	ErrIdFormat             = errors.New("id must be a valid uuid")
	ErrRequired 			= errors.New("all fields are required")
	ErrPasswordNotMatch 		= errors.New("current password is not match")
)
