package userForum

import (
	"errors"
)

var (
	ErrFailedCreateReadingList = errors.New("failed created user forum data")

	ErrRequiredId      = errors.New("id is required")
	ErrRequiredUserId  = errors.New("user id is required")
	ErrRequiredForumId = errors.New("forum id is required")
	ErrRequired        = errors.New("all fields are required")
)
