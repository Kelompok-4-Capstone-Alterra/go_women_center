package usecase

import (
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecase interface {
	GenerateTransaction() error
}

type transactionUsecase struct {
	serverKey string
}

func NewtransactionUsecase(inputServerKey string) TransactionUsecase {
	return &transactionUsecase{
		serverKey: inputServerKey,
	}
}

func (t *transactionUsecase) GenerateTransaction() error {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(t.serverKey, midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12457",
			GrossAmt: 100000,
		}, 
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Does",
			Email: "john@does.com",
			Phone: "081234567890",
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	log.Println(snapResp)
	
	return nil
}