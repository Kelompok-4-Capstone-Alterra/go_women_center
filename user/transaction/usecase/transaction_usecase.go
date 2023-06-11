package usecase

type TransactionUsecase interface {
	GenerateTransaction() error
}

type transactionUsecase struct {}

func NewtransactionUsecase() *transactionUsecase {
	return &transactionUsecase{}
}

func (t *transactionUsecase) GenerateTransaction() error {
	return nil
}