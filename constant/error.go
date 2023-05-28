package constant

import "errors"

var (
	ErrUserNotInCache = errors.New("the user isn't in cache")
	ErrExpiredCache = errors.New("the otp is already expired")
	ErrInvalidCredential = errors.New("invalid credential")
)