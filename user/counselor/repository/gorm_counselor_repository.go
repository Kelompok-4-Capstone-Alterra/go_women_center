package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(offset, limit int) ([]counselor.GetAllResponse, error)
	Count() (int, error)
	// GetById(id string) (counselor.GetByIdResponse, error)
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func (r *mysqlCounselorRepository) GetAll(offset, limit int) ([]counselor.GetAllResponse, error) {
	var counselors []counselor.GetAllResponse

	err := r.DB.Model(&entity.Counselor{}).Offset(offset).Limit(limit).Find(&counselors).Error

	if err != nil {
		return nil, err
	}

	return counselors, nil
}

func (r *mysqlCounselorRepository) Count() (int, error) {

	var totalData int64

	err := r.DB.Model(&entity.Counselor{}).Count(&totalData).Error

	if err != nil {
		return 0, err
	}

	return int(totalData), nil
}