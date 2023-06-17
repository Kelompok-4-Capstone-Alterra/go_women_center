package transaction

import "errors"

var (
	ErrorInvalidTopic            error = errors.New("invalid topic code")
	ErrorInsertDB                error = errors.New("error inserting new data to db")
	ErrorInvalidRequest          error = errors.New("error invalid callback")
	ErrorTransactionNotFound     error = errors.New("transaction not found in db")
	ErrorInvalidPaymentStatus    error = errors.New("error invalid payment status")
	ErrorMidtrans                error = errors.New("error when sending payment to midtrans")
	ErrRequired                  error = errors.New("all fields are required")
	ErrInvalidUUID               error = errors.New("invalid uuid format")
	ErrInvalidConsultationMethod error = errors.New("invalid consultation method")
	ErrDateNotFound              error = errors.New("date not found")
	ErrTimeNotFound              error = errors.New("time not found")
	ErrScheduleUnavailable       error = errors.New("schedule is unavailable")
	ErrInvalidUserCredential     error = errors.New("invalid user credential")
	ErrVoucherNotFound           error = errors.New("voucher not found")
	ErrVoucherExpired            error = errors.New("voucher unavailable")
	ErrCounselorNotFound         error = errors.New("counselor not found")
	ErrDeletingVoucher           error = errors.New("error deleting voucher")
)

var (
	// scuffed db error handling
	ErrDuplicateKey error = errors.New("duplicated key not allowed")
	ErrRecordNotFound error = errors.New("record not found")
)
