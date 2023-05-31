package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CareerRepository interface {
	GetAll(offset, limit int) ([]career.GetAllResponse, error)
	GetById(id string) (career.GetByResponse, error)
	Count() (int, error)
}

type mysqlCareerRepository struct {
	DB *gorm.DB
}

func NewMysqlCareerRepository(db *gorm.DB) CareerRepository {
	return &mysqlCareerRepository{DB: db}
}

func (r *mysqlCareerRepository) GetAll(offset, limit int) ([]career.GetAllResponse, error) {
	var career []career.GetAllResponse

	err := r.DB.Model(&entity.Career{}).Offset(offset).Limit(limit).Find(&career).Error
	if err != nil {
		return nil, err
	}
	return career, nil
}

func (r *mysqlCareerRepository) GetById(id string) (career.GetByResponse, error) {
	var career career.GetByResponse
	err := r.DB.Model(&entity.Career{}).First(&career, "id = ?", id).Error
	if err != nil {
		return career, err
	}
	return career, nil
}

func (r *mysqlCareerRepository) Count() (int, error) {
	var count int64
	err := r.DB.Model(&entity.Career{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
