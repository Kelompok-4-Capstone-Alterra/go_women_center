package readingList

import (
	"errors"
)

var (
	ErrFailedGetReadingList       = errors.New("failed to get all reading list data")
	ErrFailedGetDetailReadingList = errors.New("failed to get reading list data details")
	ErrFailedCreateReadingList    = errors.New("failed created reading list data")
	ErrFailedUpdateReadingList    = errors.New("failed to updated reading list data")
	ErrFailedDeleteReadingList    = errors.New("failed to delete reading list data")
	ErrPageNotFound               = errors.New("page not found")

	ErrRequiredId          = errors.New("id is required")
	ErrRequiredUserId      = errors.New("user id is required")
	ErrRequiredName        = errors.New("name is required")
	ErrRequiredDescription = errors.New("description is required")
	ErrRequired            = errors.New("all fields are required")
	ErrInvalidSort         = errors.New("invalid sort")
)
