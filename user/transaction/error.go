package transaction

import "errors"

var (
	ErrorInvalidGenre error = errors.New("invalid genre code")
	ErrorInsertDB error = errors.New("error inserting new data to db")
)