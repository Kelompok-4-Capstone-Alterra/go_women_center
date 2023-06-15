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
	ErrMaxOtpAttempt     = errors.New("max otp attempt")
	ErrPasswordLength    = errors.New("password must be at least 8 characters")
	ErrEmailFormat       = errors.New("email must be a valid email")
	ErrIdFormat          = errors.New("id must be a valid uuid")
	ErrRequired          = errors.New("all fields are required")
	ErrDataNotFound      = errors.New("data not found")

	// conflict
	ErrUserIsRegistered = errors.New("user is already registered")
)
