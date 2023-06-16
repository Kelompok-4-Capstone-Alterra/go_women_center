package usecase

import (
	// "log"

	"net/http"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	counselorRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	scheduleRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/repository"
	voucherRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecase interface {
	SendTransaction(transactionRequest transaction.SendTransactionRequest) (code int, res transaction.SendTransactionResponse, err error)
	UpdateStatus(transactionId string, transactionStatus string) error
	GetAll(userId string) ([]entity.Transaction, error)
	UserJoinNotification(transactionId string) error
}

type transactionUsecase struct {
	repo                 trRepo.MysqlTransactionRepository
	serverKey            string
	uuidGenerator        helper.UuidGenerator
	paymentNotifCallback string
	counselorRepo        counselorRepo.CounselorRepository
	scheduleRepo         scheduleRepo.ScheduleRepository
	voucherRepo          voucherRepo.MysqlVoucherRepository
}

func NewtransactionUsecase(
	inputServerKey string,
	uuidGenerator helper.UuidGenerator,
	trRepo trRepo.MysqlTransactionRepository,
	notifUrl string,
	cRepo            counselorRepo.CounselorRepository,
	scheduleRepo scheduleRepo.ScheduleRepository,
	voucherRepo voucherRepo.MysqlVoucherRepository,
) TransactionUsecase {
	return &transactionUsecase{
		serverKey:            inputServerKey,
		uuidGenerator:        uuidGenerator,
		repo:                 trRepo,
		paymentNotifCallback: notifUrl,
		counselorRepo: 		  cRepo,
		scheduleRepo:         scheduleRepo,
		voucherRepo:          voucherRepo,
	}
}

func (u *transactionUsecase) SendTransaction(trRequest transaction.SendTransactionRequest) (int, transaction.SendTransactionResponse, error) {
	// Initiate Snap client
	var s = snap.Client{}
	s.New(u.serverKey, midtrans.Sandbox) // sandbox

	s.Options.SetPaymentOverrideNotification(u.paymentNotifCallback)

	// generate transaction id
	transactionId, err := u.uuidGenerator.GenerateUUID()
	if err != nil {
		return http.StatusInternalServerError,
			transaction.SendTransactionResponse{},
			err
	}

	res := transaction.SendTransactionResponse{}

	// check topic availability
	trTopic, ok := constant.TOPICS[trRequest.CounselorTopicKey]
	if !ok {
		return http.StatusBadRequest,
			transaction.SendTransactionResponse{},
			transaction.ErrorInvalidTopic
	}

	// check date availability
	_, err = u.scheduleRepo.GetDateById(trRequest.ConsultationDateID)
	if err != nil {
		return http.StatusBadRequest,
			transaction.SendTransactionResponse{},
			transaction.ErrDateNotFound
	}

	// check time availability
	_, err = u.scheduleRepo.GetTimeById(trRequest.ConsultationTimeID)
	if err != nil {
		return http.StatusBadRequest,
			transaction.SendTransactionResponse{},
			transaction.ErrTimeNotFound
	}

	// implement voucher
	voucherValue := int64(0)
	if trRequest.VoucherId != "" {
		
		// TODO: check voucher availability
		savedVoucher, err := u.voucherRepo.GetById(trRequest.UserCredential.ID, trRequest.VoucherId)
		if err != nil {
			return http.StatusBadRequest,
			transaction.SendTransactionResponse{},
			transaction.ErrVoucherNotFound
		}

		// TODO: validate voucher exp date
		timeNow := time.Now()
		if timeNow.After(savedVoucher.ExpDate) {
			return http.StatusBadRequest,
			transaction.SendTransactionResponse{},
			transaction.ErrVoucherExpired
		}

		voucherValue = savedVoucher.Value
	}


	// TODO: get gross price from counselor data
	savedCounselor, err := u.counselorRepo.GetById(trRequest.CounselorID)	
	if err != nil {
		return http.StatusBadRequest,
		transaction.SendTransactionResponse{},
		transaction.ErrCounselorNotFound
	}

	// TODO: calculate total price
	totalPrice := int64(savedCounselor.Price) - voucherValue
	if totalPrice < 0 {
		totalPrice = 0
	}

	// initialize db data model
	transactionData := entity.Transaction{
		ID:                 transactionId,
		UserId:             trRequest.UserCredential.ID,
		DateId:             trRequest.ConsultationDateID,
		CounselorID:        trRequest.CounselorID,
		CounselorTopic:     trTopic[0],
		TimeId:             trRequest.ConsultationTimeID,
		TimeStart:          trRequest.ConsultationTimeStart,
		ConsultationMethod: trRequest.ConsultationMethod,
		Status:             "pending",
		ValueVoucher:       voucherValue,
		GrossPrice:         int64(savedCounselor.Price),
		TotalPrice:         totalPrice,
		IsReviewed:         false,
		Created_at:         time.Now(),
	}

	_, err = u.repo.CreateTransaction(transactionData)
	if err != nil {
		if err.Error() == transaction.ErrDuplicateKey.Error() {
			return http.StatusBadRequest,
				transaction.SendTransactionResponse{},
				transaction.ErrScheduleUnavailable
		}
		return http.StatusInternalServerError,
			transaction.SendTransactionResponse{},
			err
	}

	// Initiate Snap request param
	// using total price as grossamt
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionId,
			GrossAmt: totalPrice,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: trRequest.UserCredential.Username,
			Email: trRequest.UserCredential.Email,
		},
	}

	// Execute request create Snap transaction to Midtrans Snap API
	snapResp, snapErr := s.CreateTransaction(req)
	if snapErr != nil {
		return http.StatusInternalServerError,
			transaction.SendTransactionResponse{},
			transaction.ErrorMidtrans
	}

	res.TransactionID = transactionId
	res.PaymentLink = snapResp.RedirectURL
	res.Data = transactionData

	// TODO: delete voucher if implemented
	_, err = u.voucherRepo.DeleteById(trRequest.UserCredential.ID, trRequest.VoucherId)
	if err != nil {
		return http.StatusInternalServerError,
			transaction.SendTransactionResponse{},
			transaction.ErrDeletingVoucher
	}

	return http.StatusOK, res, nil
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

// user join the consultation
func (u *transactionUsecase) UserJoinNotification(transactionId string) error {
	err := u.repo.UpdateStatusById(transactionId, "completed")
	// TODO: better error handling
	if err != nil {
		return err
	}
	return nil
}
