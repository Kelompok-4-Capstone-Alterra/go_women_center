package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type MysqlTransactionRepository interface {
	GetAll() ([]entity.Transaction, error)
}

type mysqlTransactionRepository struct {
	DB *gorm.DB
}

func NewMysqltransactionRepository(db *gorm.DB) MysqlTransactionRepository {
	return &mysqlTransactionRepository{
		DB: db,
	}
}

func (tr *mysqlTransactionRepository) GetAll() ([]entity.Transaction, error) {
	transactionData := []entity.Transaction{}
	err := tr.DB.Find(&transactionData).Error
	if err != nil {
		return nil, err
	}
	return transactionData, nil
}