package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"gorm.io/gorm"
)

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) domain.CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func (r *mysqlCounselorRepository) GetAll(offset, limit int) ([]domain.Counselor, error) {
	var counselor []domain.Counselor

	err := r.DB.Offset(offset).Limit(limit).Find(&counselor).Error
	if err != nil {
		return nil, err
	}
	return counselor, nil
}

func (r *mysqlCounselorRepository) Count() (int, error) {
	var count int64
	err := r.DB.Model(&domain.Counselor{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *mysqlCounselorRepository) GetById(id string) (domain.Counselor, error) {
	var counselor domain.Counselor
	err := r.DB.Where("id = ?", id).Find(&counselor).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func (r *mysqlCounselorRepository) Create(counselor domain.Counselor) error {
	err := r.DB.Create(&counselor).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlCounselorRepository) Update(id string, counselor domain.Counselor) error {
	return nil
}

func (r *mysqlCounselorRepository) Delete(id string) error {
	return nil
}

