package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	trError "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	"gorm.io/gorm"
)

type MysqlTransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
}

type mysqlTransactionRepository struct {
	DB *gorm.DB
}

func NewMysqltransactionRepository(db *gorm.DB) MysqlTransactionRepository {
	return &mysqlTransactionRepository{
		DB: db,
	}
}

func (tr *mysqlTransactionRepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	err := tr.DB.Create(&transaction).Error
	if err != nil {
		return entity.Transaction{}, trError.ErrorInsertDB
	}
	return transaction, nil
}
