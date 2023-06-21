package transaction

import "errors"

var (
	ErrUpdate = errors.New("transaction not found/link already sent")
)

// gorm error
var (
	ErrEmptySlice = errors.New("empty slice found")
)


// error validation
var (
	ErrPage = errors.New("invalid page format")
	ErrLimit = errors.New("invalid limit format")
	ErrSearch = errors.New("invalid search format")
	ErrSortBy = errors.New("invalid sort by format")
	ErrTransactionId = errors.New("invalid transaction id format")
	ErrInvalidLink = errors.New("invalid link format")
	ErrInvalidStartDate = errors.New("invalid start date")
	ErrInvalidEndDate = errors.New("invalid end date")
)