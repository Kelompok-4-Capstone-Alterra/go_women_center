package repository

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type MysqlVoucherRepository interface {
	GetAll(userId string) ([]entity.Voucher, error)
	GetById(userId, voucherId string) (entity.Voucher, error)
	DeleteById(userId, voucherId string) (entity.Voucher, error)
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
	timeNow := time.Now()
	allUserVoucher := []entity.Voucher{}
	err := tr.DB.Where("user_id = ? AND exp_date > ?", userId, timeNow).Find(&allUserVoucher).Error
	if err != nil {
		return nil, err
	}

	return allUserVoucher, nil
}

func (tr *mysqlVoucherRepository) GetById(userId, voucherId string) (entity.Voucher, error) {
	userVoucher := entity.Voucher{}
	err := tr.DB.Where("user_id = ? AND id = ?", userId, voucherId).Find(&userVoucher).Error
	if err != nil {
		return entity.Voucher{}, err
	}

	return userVoucher, nil
}

func (tr *mysqlVoucherRepository) DeleteById(userId, voucherId string) (entity.Voucher, error) {
	userVoucher := entity.Voucher{}
	err := tr.DB.Where("user_id = ? AND id = ?", userId, voucherId).Delete(&userVoucher).Error
	if err != nil {
		return entity.Voucher{}, err
	}

	return userVoucher, nil
}
