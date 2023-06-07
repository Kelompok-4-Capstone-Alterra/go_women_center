package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type DateRepository interface {
	GetByCounselorId(counselorId string) ([]entity.Date, error)
	Create(date entity.Date) error
	DeleteByCounselorId(counselorId string) error
}

type mysqlDateRepository struct {
	DB *gorm.DB
}


func NewMysqlDateRepository(db *gorm.DB) DateRepository {
	return &mysqlDateRepository{db}
}

func(r *mysqlDateRepository) GetByCounselorId(counselorId string) ([]entity.Date, error) {

	var date []entity.Date
	err := r.DB.Find(&date, "counselor_id = ?", counselorId).Error

	if err != nil {
		return nil, err
	}
	return date, nil
}

func (r *mysqlDateRepository) Create(date entity.Date) error {

	err := r.DB.Create(&date).Error

	if err != nil {
		return err
	}

	return nil
}

func(r *mysqlDateRepository) DeleteByCounselorId(counselorId string) error {
	
	err := r.DB.Delete(&entity.Date{}, "counselor_id = ?", counselorId).Error
	
	if err != nil {
		return err
	}

	return nil
}

