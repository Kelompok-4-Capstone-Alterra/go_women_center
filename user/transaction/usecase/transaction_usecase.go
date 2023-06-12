package usecase

import (
	// "log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecase interface {
	GenerateTransaction(transactionRequest transaction.GenerateTransactionRequest) (transaction.GenerateTransactionResponse, error)
}

type transactionUsecase struct {
	serverKey string
	uuidGenerator helper.UuidGenerator
}

func NewtransactionUsecase(
		inputServerKey string,
		uuidGenerator helper.UuidGenerator,
	) TransactionUsecase {
	return &transactionUsecase{
		serverKey: inputServerKey,
		uuidGenerator: uuidGenerator,
	}
}

func (t *transactionUsecase) GenerateTransaction(transactionRequest transaction.GenerateTransactionRequest) (transaction.GenerateTransactionResponse, error) {
	res := transaction.GenerateTransactionResponse{}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(t.serverKey, midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	transactionId, err := t.uuidGenerator.GenerateUUID()
	if err != nil {
		return res, err
	}

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionId,
			GrossAmt: 100000,
		}, 
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: transactionRequest.Email,
			Phone: "081234567890",
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	
	// TODO: send order id
	snapResp, _ := s.CreateTransaction(req)

	res.ID = transactionId
	res.Link = snapResp.RedirectURL

	// TODO: create transaction in db with status pending
	// catch callback res from midtrans
	// if status 200 then update status success
	// else then update status to canceled
	
	return res, nil
}

// success only
func getAll() {}

// check status after payment
func verifyById(id string) {}