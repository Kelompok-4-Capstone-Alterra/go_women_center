package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(offset, limit int) ([]counselor.GetAllResponse, error)
	Count() (int, error)
	GetByEmail(email string) (counselor.GetByResponse, error)
	GetById(id string) (counselor.GetByResponse, error)
	Create(counselor entity.Counselor) error
	Update(id string, counselor entity.Counselor) error
	Delete(id string) error
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func (r *mysqlCounselorRepository) GetAll(offset, limit int) ([]counselor.GetAllResponse, error) {
	var counselor []counselor.GetAllResponse
	
	err := r.DB.Model(&entity.Counselor{}).Offset(offset).Limit(limit).Find(&counselor).Error
	if err != nil {
		return nil, err
	}
	return counselor, nil
}

func (r *mysqlCounselorRepository) Count() (int, error) {
	var count int64
	err := r.DB.Model(&entity.Counselor{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *mysqlCounselorRepository) GetById(id string) (counselor.GetByResponse, error) {
	var counselor counselor.GetByResponse
	err := r.DB.Model(&entity.Counselor{}).First(&counselor, "id = ?", id).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func (r *mysqlCounselorRepository) GetByEmail(email string) (counselor.GetByResponse, error) {
	var counselor counselor.GetByResponse
	err := r.DB.Model(&entity.Counselor{}).First(&counselor, "email = ?", email).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func (r *mysqlCounselorRepository) Create(counselor entity.Counselor) error {
	err := r.DB.Create(&counselor).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlCounselorRepository) Update(id string, counselor entity.Counselor) error {
	
	err := r.DB.Model(&entity.Counselor{}).Where("id = ?", id).Updates(counselor).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlCounselorRepository) Delete(id string) error {

	err := r.DB.Delete(&entity.Counselor{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

