package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type MysqlTransactionRepository interface {
	GetAll(search, sortBy string, offset, limit int) ([]entity.Transaction, int64, error)
	GetById(id string) (entity.Transaction, error)
	UpdateById(id, link, status string) (entity.Transaction, error)
}

type mysqlTransactionRepository struct {
	DB *gorm.DB
}

func NewMysqltransactionRepository(db *gorm.DB) MysqlTransactionRepository {
	return &mysqlTransactionRepository{
		DB: db,
	}
}

func (tr *mysqlTransactionRepository) GetAll(search, sortBy string, offset, limit int) ([]entity.Transaction, int64, error) {
	transactionData := []entity.Transaction{}
	count := int64(0)
	err := tr.DB.
		Debug().
		Model(&entity.Transaction{}).
		Where(
			"counselor_topic LIKE ? OR consultation_method LIKE ? OR date_id LIKE ? OR time_id LIKE ? OR id LIKE ? OR user_id LIKE ? OR counselor_id LIKE ? OR status LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Order(sortBy).
		Find(&transactionData).Error

	if err != nil {
		return nil, 0, err
	}

	return transactionData, count, nil
}

func (tr *mysqlTransactionRepository) GetById(id string) (entity.Transaction, error) {
	transactionData := entity.Transaction{}
	err := tr.DB.Where("id = ?", id).First(&transactionData).Error
	if err != nil {
		return entity.Transaction{}, err
	}
	return transactionData, nil
}

func (tr *mysqlTransactionRepository) UpdateById(id, link, status string) (entity.Transaction, error) {
	updatedData := entity.Transaction{
		Status: status,
		Link:   link,
	}

	result := tr.DB.Debug().
		Model(&entity.Transaction{}).
		Where("id = ?", id).
		Updates(&updatedData)

	updated := result.RowsAffected
	if updated < 1 {
		return entity.Transaction{}, gorm.ErrEmptySlice
	}

	err := result.Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return updatedData, nil
}
