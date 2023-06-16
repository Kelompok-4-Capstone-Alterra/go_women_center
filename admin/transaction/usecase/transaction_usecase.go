package usecase

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type TransactionUsecase interface {
	GetAll() (code int, data []entity.Transaction, err error)
	SendLink(req transaction.SendLinkRequest) (code int, err error)
}

type transactionUsecase struct {
	repo trRepo.MysqlTransactionRepository
}

func NewtransactionUsecase(
	trRepo trRepo.MysqlTransactionRepository,
) TransactionUsecase {
	return &transactionUsecase{
		repo: trRepo,
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
	err := tu.repo.UpdateById(req.TransactionId, req.Link)
	if err != nil {
		if err.Error() == transaction.ErrEmptySlice.Error() {
			return http.StatusBadRequest, transaction.ErrUpdate
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
