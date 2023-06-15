package usecase

import (
	// "log"

	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	counselorUC "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecase interface {
	SendTransaction(transactionRequest transaction.SendTransactionRequest) (transaction.SendTransactionResponse, error)
	UpdateStatus(transactionId string, transactionStatus string) error
	GetAll(userId string) ([]entity.Transaction, error)
}

type transactionUsecase struct {
	repo                 trRepo.MysqlTransactionRepository
	serverKey            string
	uuidGenerator        helper.UuidGenerator
	Counselor            counselorUC.CounselorUsecase
	paymentNotifCallback string
}

func NewtransactionUsecase(
	inputServerKey string,
	uuidGenerator helper.UuidGenerator,
	trRepo trRepo.MysqlTransactionRepository,
	notifUrl string,
) TransactionUsecase {
	return &transactionUsecase{
		serverKey:            inputServerKey,
		uuidGenerator:        uuidGenerator,
		repo:                 trRepo,
		paymentNotifCallback: notifUrl,
	}
}

func (u *transactionUsecase) SendTransaction(trRequest transaction.SendTransactionRequest) (transaction.SendTransactionResponse, error) {
	// Initiate Snap client
	var s = snap.Client{}
	s.New(u.serverKey, midtrans.Sandbox) // sandbox

	s.Options.SetPaymentOverrideNotification(u.paymentNotifCallback)

	// generate transaction id
	transactionId, err := u.uuidGenerator.GenerateUUID()
	if err != nil {
		return transaction.SendTransactionResponse{}, err
	}

	// Initiate Snap request param
	// using total price as grossamt
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionId,
			GrossAmt: trRequest.TotalPrice,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: trRequest.UserCredential.Username,
			Email: trRequest.UserCredential.Email,
		},
	}

	res := transaction.SendTransactionResponse{}

	// check topic availability
	trTopic, ok := constant.TOPICS[trRequest.CounselorTopicKey]
	if !ok {
		return transaction.SendTransactionResponse{}, transaction.ErrorInvalidGenre
	}

	// initialize db data model
	transactionData := entity.Transaction{
		ID:                 transactionId,
		UserId:             trRequest.UserCredential.ID,
		DateId:             trRequest.CounselingDateID,
		CounselorId:        trRequest.CounselorID,
		CounselorTopic:     trTopic[0],
		TimeId:             trRequest.CounselingTimeID,
		TimeStart:          trRequest.CounselingTimeStart,
		ConsultationMethod: trRequest.CounselingMethod,
		Status:             "pending",
		ValueVoucher:       trRequest.ValueVoucher,
		GrossPrice:         trRequest.GrossPrice,
		TotalPrice:         trRequest.TotalPrice,
		IsReviewed:         false,
		Created_at:         time.Now(),
	}

	data, err := u.repo.CreateTransaction(transactionData)
	if err != nil {
		return transaction.SendTransactionResponse{}, err
	}

	// Execute request create Snap transaction to Midtrans Snap API
	snapResp, snapErr := s.CreateTransaction(req)
	if snapErr != nil {
		return transaction.SendTransactionResponse{}, transaction.ErrorMidtrans
	}

	res.TransactionID = transactionId
	res.PaymentLink = snapResp.RedirectURL

	res.Data = data

	return res, nil
}

/*
catch callback res from midtrans

if status 200 then update status success

else then status is the same
*/
func (u *transactionUsecase) UpdateStatus(transactionId string, transactionStatus string) error {
	if transactionStatus != "settlement" {
		return nil
	}

	transactionStatus = "ongoing"

	savedTransaction, err := u.verifyById(transactionId)
	if err != nil {
		return err
	}

	_, err = u.repo.UpdateStatusByData(savedTransaction, transactionStatus)
	if err != nil {
		return transaction.ErrorInsertDB
	}

	return nil
}

// check status after payment
func (u *transactionUsecase) verifyById(id string) (entity.Transaction, error) {
	savedTransaction, err := u.repo.GetById(id)
	if err != nil {
		return entity.Transaction{}, err
	}
	return savedTransaction, nil
}

// success only
func (u *transactionUsecase) GetAll(userId string) ([]entity.Transaction, error) {
	data, err := u.repo.GetAllSuccess(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
