package repository

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type MysqlTransactionRepository interface {
	GetAll(search, sortBy string, offset, limit int) ([]entity.Transaction, int64, error)
	GetById(id string) (entity.Transaction, error)
	UpdateById(id, link, status string) (entity.Transaction, error)
	GetAllForReport(transaction.ReportRequest) ([]entity.Transaction, int64, error)
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
		Preload("Counselor").
		Joins("left join counselors on counselors.id = transactions.counselor_id").
		Where("transactions.status != ?", "pending").
		Where(
			"transactions.consultation_method LIKE ? OR transactions.id LIKE ? OR transactions.user_id LIKE ? OR transactions.counselor_id LIKE ? OR transactions.status LIKE ? OR counselors.username LIKE ?",
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
		log.Println(err)
		return nil, 0, err
	}

	log.Println(transactionData)

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

// if func is called for displaying in web, will use pagination data
// 
// else will not be using pagination
func (tr *mysqlTransactionRepository) GetAllForReport(tReq transaction.ReportRequest) ([]entity.Transaction, int64, error) {
	search := tReq.Search
	sortBy := tReq.SortBy
	
	transactionData := []entity.Transaction{}
	count := int64(0)
	
	dbQuery := tr.DB.
		Debug().
		Model(&entity.Transaction{}).
		Preload("Counselor").
		Joins("left join counselors on counselors.id = transactions.counselor_id").
		Where("transactions.status != ?", "pending").
		Where(
			"transactions.consultation_method LIKE ? OR transactions.id LIKE ? OR transactions.user_id LIKE ? OR transactions.counselor_id LIKE ? OR transactions.status LIKE ? OR counselors.username LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%")

	log.Println(transactionData)

	if tReq.StartDate != "" && tReq.EndDate != "" {
		dbQuery.
			Where("transactions.created_at BETWEEN ? AND ?", tReq.StartDate, tReq.EndDate)
	}

	if tReq.StartDate != "" && tReq.EndDate == "" {
		dbQuery.
			Where("transactions.created_at >= ?", tReq.StartDate)
	}

	if tReq.StartDate == "" && tReq.EndDate != "" {
		dbQuery.
			Where("transactions.created_at <= ?", tReq.EndDate)
	}

	// for pagination
	if !tReq.IsDownload {
		dbQuery.
			Count(&count).
			Offset(tReq.Offset).
			Limit(tReq.Limit)
	}

	dbQuery.
		Order(sortBy).
		Find(&transactionData)

	log.Println(transactionData)

	if err := dbQuery.Error; err != nil {
		return nil, 0, err
	}

	return transactionData, count, nil
}
