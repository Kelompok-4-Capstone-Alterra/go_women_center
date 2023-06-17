package usecase

import (
	"net/http"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type TransactionUsecase interface {
	GetAll() (code int, data []entity.Transaction, err error)
	SendLink(req transaction.SendLinkRequest) (code int, err error)
	CancelTransaction(req transaction.CancelTransactionRequest) (int, error)
}

type transactionUsecase struct {
	repo        trRepo.MysqlTransactionRepository
	voucherRepo trRepo.MysqlVoucherRepository
	uuidGen     helper.UuidGenerator
}

func NewtransactionUsecase(
	trRepo trRepo.MysqlTransactionRepository,
	voucherRepo trRepo.MysqlVoucherRepository,
	uuidGen helper.UuidGenerator,
) TransactionUsecase {
	return &transactionUsecase{
		repo:        trRepo,
		voucherRepo: voucherRepo,
		uuidGen:     uuidGen,
	}
}

func (tu *transactionUsecase) GetAll() (int, []entity.Transaction, error) {
	// TODO: pagination
	data, err := tu.repo.GetAll()
	if err != nil {
		return http.StatusInternalServerError,
			nil,
			err
	}
	return http.StatusOK,
		data,
		nil
}

func (tu *transactionUsecase) SendLink(req transaction.SendLinkRequest) (int, error) {
	_, err := tu.repo.UpdateById(req.TransactionId, req.Link, "waiting")
	if err != nil {
		if err.Error() == transaction.ErrEmptySlice.Error() {
			return http.StatusBadRequest, transaction.ErrUpdate
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (tu *transactionUsecase) CancelTransaction(req transaction.CancelTransactionRequest) (int, error) {
	// TODO: implement rollback
	transactionData, err := tu.repo.GetById(req.TransactionId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = tu.repo.UpdateById(req.TransactionId, "-", "canceled")
	if err != nil {
		if err.Error() == transaction.ErrEmptySlice.Error() {
			return http.StatusBadRequest, transaction.ErrUpdate
		}
		return http.StatusInternalServerError, err
	}

	voucherId, err := tu.uuidGen.GenerateUUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	expDate := time.Now().AddDate(0, 1, 0)
	voucher := entity.Voucher{
		ID:      voucherId,
		UserId:  transactionData.UserId,
		Value:   transactionData.GrossPrice,
		ExpDate: expDate,
	}

	_, err = tu.voucherRepo.CreateVoucher(voucher)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
