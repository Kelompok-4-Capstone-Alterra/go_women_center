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
)
