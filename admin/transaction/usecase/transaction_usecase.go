package usecase

import (
	"net/http"

	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type TransactionUsecase interface {
	GetAll() (code int, data []entity.Transaction, err error)
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