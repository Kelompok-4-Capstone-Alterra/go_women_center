package readingListArticle

import (
	"errors"
)

var (
	ErrFailedAddReadingListArticle    = errors.New("failed to join reading list article")
	ErrFailedDeleteReadingListArticle = errors.New("failed to delete reading lists artcile")
	ErrPageNotFound                   = errors.New("page not found")
)
