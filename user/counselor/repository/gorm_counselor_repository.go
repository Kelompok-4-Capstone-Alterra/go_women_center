package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(search, topic, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error)
	GetById(id string) (counselor.GetByResponse, error)
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func(r *mysqlCounselorRepository) GetAll(search, topic, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error) {

	var counselors []counselor.GetAllResponse
	var totalData int64
	err := r.DB.Model(&entity.Counselor{}).
		Where("topic = ? AND name LIKE ?", topic, "%"+search+"%").
		Count(&totalData).
		Order(sortBy).
		Offset(offset).
		Limit(limit).
		Find(&counselors).Error

	if err != nil {
		return nil, totalData, err
	}

	return counselors, totalData, nil
}

func(r *mysqlCounselorRepository) GetById(id string) (counselor.GetByResponse, error) {
	
	var counselor counselor.GetByResponse

	err := r.DB.Model(&entity.Counselor{}).Where("id = ?", id).First(&counselor).Error
	
	if err != nil {
		return counselor, err
	}

	return counselor, nil
}