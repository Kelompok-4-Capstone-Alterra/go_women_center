package repository

import (
	"log"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"

	"gorm.io/gorm"
)

type MysqlTransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	GetAll(userId, search string, trStatus... string) ([]entity.Transaction, error)
	GetById(id string) (entity.Transaction, error)
	UpdateStatusByData(savedData entity.Transaction, newStatus string) (entity.Transaction, error)
	UpdateStatusById(id string, newStatus string) error
	GetOccurTransacTodayByCounselorId(counselorId string) ([]entity.Transaction, error)
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
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (tr *mysqlTransactionRepository) GetAll(userId, search string, trStatus... string) ([]entity.Transaction, error) {
	allUserTransaction := []entity.Transaction{}
	gormQuery := tr.DB.Debug().Preload("Counselor")

	log.Println(trStatus[0])
	
	if len(trStatus) > 1 {
		gormQuery.
			Where("user_id = ?", userId).
			Where("status = ? OR status = ?", trStatus[0], trStatus[1])
	} else {
		gormQuery.
			Where("user_id = ? AND status = ?", userId, trStatus[0])
	}

	if search != "" {
		gormQuery.Where(
			"counselor_topic LIKE ? OR consultation_method LIKE ? OR date_id LIKE ? OR time_id LIKE ? OR id LIKE ? OR counselor_id LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%")
	}

	err := gormQuery.
		Find(&allUserTransaction).
		Error
	if err != nil {
		return nil, err
	}

	return allUserTransaction, nil
}

func (tr *mysqlTransactionRepository) GetById(id string) (entity.Transaction, error) {
	savedTransaction := entity.Transaction{}
	err := tr.DB.Where("id = ?", id).First(&savedTransaction).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return savedTransaction, nil
}

func (tr *mysqlTransactionRepository) UpdateStatusByData(savedData entity.Transaction, newStatus string) (entity.Transaction, error) {
	savedData.Status = newStatus
	err := tr.DB.Updates(&savedData).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return savedData, nil
}

func (tr *mysqlTransactionRepository) UpdateStatusById(id string, newStatus string) error {
	result := tr.DB.Debug().Model(&entity.Transaction{}).Where("id = ?", id).Update("status", newStatus)

	updated := result.RowsAffected
	if updated < 1 {
		return gorm.ErrEmptySlice
	}

	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

// get all transaction by counselor id for today
func (tr *mysqlTransactionRepository) GetOccurTransacTodayByCounselorId(counselorId string) ([]entity.Transaction, error) {
	currentTime := time.Now()
	currentDate := currentTime.Format(time.DateOnly)

	var transactions []entity.Transaction
	err := tr.DB.Model(&entity.Transaction{}).Where("counselor_id = ? AND DATE(created_at) = ?", counselorId, currentDate).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

/*
create transaction db data, return db instance with transaction

call CommitCreate if payment success

call RollbackCreate if payment failed
*/
func (tr *mysqlTransactionRepository) StartCreate(transaction entity.Transaction) (*gorm.DB, error) {
	tx := tr.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		return nil, err
	}

	err = tx.Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// cancel creating transaction data
// if sending payment req to payment gateway failed
func (tr *mysqlTransactionRepository) RollbackCreate(tx *gorm.DB) {
	tx.Rollback()
}

// commit creating transaction data
// if sending payment req to payment gateway success
func (tr *mysqlTransactionRepository) CommitCreate(tx *gorm.DB) error {
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
