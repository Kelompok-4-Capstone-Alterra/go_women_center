package userForum

import (
	"errors"
)

var (
	ErrFailedGetDetailUserForum = errors.New("failed to get member forum data details")
	ErrFailedCreateUserForum    = errors.New("failed created user forum data")
	ErrFailedJoinUserForum      = errors.New("failed to join forum")

	ErrRequiredId      = errors.New("id is required")
	ErrRequiredUserId  = errors.New("user id is required")
	ErrRequiredForumId = errors.New("forum id is required")
	ErrRequired        = errors.New("all fields are required")
)
