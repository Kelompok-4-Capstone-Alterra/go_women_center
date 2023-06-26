package forum

import (
	"errors"
)

var (
	ErrFailedGetForum       = errors.New("failed to get all forum data")
	ErrFailedGetDetailForum = errors.New("failed to get forum data details")
	ErrFailedCreateForum    = errors.New("failed created forum data")
	ErrFailedUpdateForum    = errors.New("failed to updated forum data")
	ErrFailedDeleteForum    = errors.New("failed to delete forum data")
	ErrPageNotFound         = errors.New("page not found")

	ErrRequiredCategory = errors.New("category is required")
	ErrRequiredLink     = errors.New("link is required")
	ErrRequiredTopic    = errors.New("topic are required")
	ErrRequired         = errors.New("topic are required")

	ErrInvalidSort     = errors.New("invalid sorting")
	ErrInvalidMyForum  = errors.New("invalid my forum")
	ErrInvalidCategory = errors.New("invalid category forum")
	ErrInvalidId       = errors.New("invalid id forum")
	ErrNotAccess       = errors.New("cannot access this data")
	ErrInvalidLink     = errors.New("invalid link forum")
	ErrInvalidUrlHost  = errors.New("invalid url host, use telegram link")
)
