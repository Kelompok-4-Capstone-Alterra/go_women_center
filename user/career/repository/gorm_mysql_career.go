package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career"
	"gorm.io/gorm"
)

type CareerRepository interface {
	GetAll(search, sortBy  string) ([]career.GetAllResponse, error)
	GetById(id string) (entity.Career, error)
	GetBySearch(search string) ([]career.GetAllResponse, error)
	Count() (int, error)
}

type mysqlCareerRepository struct {
	DB *gorm.DB
}

func NewMysqlCareerRepository(db *gorm.DB) CareerRepository {
	return &mysqlCareerRepository{DB: db}
}

func (r *mysqlCareerRepository) GetAll(search, sortBy  string) ([]career.GetAllResponse, error) {
	//added sort
	var career []career.GetAllResponse
	var count int64
	err := r.DB.Model(&entity.Career{}).
		Where("job_position LIKE ? OR company_name LIKE ? OR Location LIKE ? OR CAST(Salary AS CHAR) LIKE ? OR company_email LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&count).
		Order(sortBy).
		Find(&career).Error

	if err != nil {
		return nil, err
	}
	return career, nil
}

func (r *mysqlCareerRepository) GetById(id string) (entity.Career, error) {
	var career entity.Career
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
