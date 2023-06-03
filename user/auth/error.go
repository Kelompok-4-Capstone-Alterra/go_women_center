package auth

import "errors"

var (
	ErrInvalidOtp        = errors.New("invalid otp")
	ErrExpiredOtp        = errors.New("the otp is already expired")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrInvalidInput      = errors.New("invalid input")
	ErrFailedEncrpyt     = errors.New("failed to encrypt credential")
)
