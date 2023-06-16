package transaction

import "errors"

var (
	ErrUpdate = errors.New("transaction not found/link already sent")
)

// gorm error
var (
	ErrEmptySlice = errors.New("empty slice found")
)
