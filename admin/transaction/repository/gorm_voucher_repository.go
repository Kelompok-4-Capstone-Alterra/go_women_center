package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlVoucherRepository interface {
	CreateVoucher(voucher entity.Voucher) (entity.Voucher, error)
}

type mysqlVoucherRepository struct {
	DB *gorm.DB
}

func NewMysqlVoucherRepository(db *gorm.DB) MysqlVoucherRepository {
	return &mysqlVoucherRepository{
		DB: db,
	}
}

func (vr *mysqlVoucherRepository) CreateVoucher(voucher entity.Voucher) (entity.Voucher, error) {
	err := vr.DB.Clauses(clause.Returning{}).Create(&voucher).Error
	if err != nil {
		return entity.Voucher{}, err
	}
	return voucher, nil
}