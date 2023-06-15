package transaction

import "errors"

var (
	ErrorInvalidGenre error = errors.New("invalid genre code")
	ErrorInsertDB error = errors.New("error inserting new data to db")
	ErrorInvalidRequest error = errors.New("error invalid callback")
	ErrorTransactionNotFound error = errors.New("transaction not found in db")
	ErrorInvalidPaymentStatus error = errors.New("error invalid payment status")
	ErrorMidtrans error = errors.New("error when sending payment to midtrans")
)