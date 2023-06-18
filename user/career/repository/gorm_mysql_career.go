package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career"
	"gorm.io/gorm"
)

type CareerRepository interface {
	GetAll(search string, offset, limit int) ([]career.GetAllResponse, int64, error)
	GetById(id string) (career.GetByResponse, error)
	GetBySearch(search string) ([]career.GetAllResponse, error)
	Count() (int, error)
}

type mysqlCareerRepository struct {
	DB *gorm.DB
}

func NewMysqlCareerRepository(db *gorm.DB) CareerRepository {
	return &mysqlCareerRepository{DB: db}
}

func (r *mysqlCareerRepository) GetAll(search string, offset, limit int) ([]career.GetAllResponse, int64, error) {
	var career []career.GetAllResponse
	var count int64
	err := r.DB.Model(&entity.Career{}).
		Where("job_position LIKE ? OR company_name LIKE ? OR Location LIKE ? OR CAST(Salary AS CHAR) LIKE ? OR company_email LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&career).Error

	if err != nil {
		return nil, 0, err
	}
	return career, count, nil
}

func (r *mysqlCareerRepository) GetById(id string) (career.GetByResponse, error) {
	var career career.GetByResponse
	err := r.DB.Model(&entity.Career{}).First(&career, "id = ?", id).Error
	if err != nil {
		return career, err
	}
	return career, nil
}

func (r *mysqlCareerRepository) GetBySearch(search string) ([]career.GetAllResponse, error) {
	var career []career.GetAllResponse

	err := r.DB.Model(&entity.Career{}).
		Where("job_position LIKE ? OR company_name LIKE ? OR Location LIKE ? OR CAST(Salary AS CHAR) LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Find(&career).Error
	if err != nil {
		return nil, err
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
