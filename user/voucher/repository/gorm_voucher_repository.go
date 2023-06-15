package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type MysqlVoucherRepository interface {
	GetAll(userId string) ([]entity.Voucher, error)
}

type mysqlVoucherRepository struct {
	DB *gorm.DB
}

func NewMysqltransactionRepository(db *gorm.DB) MysqlVoucherRepository {
	return &mysqlVoucherRepository{
		DB: db,
	}
}

func (tr *mysqlVoucherRepository) GetAll(userId string) ([]entity.Voucher, error) {
	allUserVoucher := []entity.Voucher{}
	err := tr.DB.Where("user_id = ?", userId).Find(&allUserVoucher).Error
	if err != nil {
		return nil, err
	}

	return allUserVoucher, nil
}