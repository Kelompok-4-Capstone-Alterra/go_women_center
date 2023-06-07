package auth

import "errors"

var (

	// internal server error
	ErrInternalServerError = errors.New("internal server error")

	ErrInvalidOtp        = errors.New("invalid otp")
	ErrExpiredOtp        = errors.New("the otp is already expired")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrInvalidInput      = errors.New("invalid input")
	ErrFailedEncrpyt     = errors.New("failed to encrypt credential")

	// conflict
	ErrUserIsRegistered = errors.New("user is already registered")
)
