package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type TimeRepository interface {
	GetByCounselorId(counseorId string) ([]entity.Time, error)
	Create(time entity.Time) error
	DeleteByCounselorId(counselorId string) error
}

type mysqlTimeRepository struct {
	DB *gorm.DB
}

func NewMysqlTimeRepository(db *gorm.DB) TimeRepository {
	return &mysqlTimeRepository{DB: db}
}

func(r *mysqlTimeRepository) GetByCounselorId(counselorId string) ([]entity.Time, error) {
	var times []entity.Time

	err := r.DB.Find(&times, "counselor_id = ?", counselorId).Error

	if err != nil {
		return nil, err
	}

	return times, nil
}


func(r *mysqlTimeRepository) Create(time entity.Time) error {
	
	err := r.DB.Create(&time).Error
	
	if err != nil {
		return err
	}

	return nil
}

func(r *mysqlTimeRepository) DeleteByCounselorId(counselorId string) error {
	
	err := r.DB.Delete(&entity.Time{}, "counselor_id = ?", counselorId).Error
	
	if err != nil {
		return err
	}

	return nil
}