package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type TimeRepository interface {
	GetByDateId(id string) (entity.Time, error)
	Create(Time entity.Time) error
	Update(Time entity.Time) error
	Delete(id string) error
}

type mysqlTimeRepository struct {
	DB *gorm.DB
}

func NewMysqlTimeRepository(db *gorm.DB) TimeRepository {
	return &mysqlTimeRepository{DB: db}
}

func(r *mysqlTimeRepository) GetByDateId(id string) (entity.Time, error) {

	var time entity.Time

	err := r.DB.First(&time, "date_id = ?", id).Error

	if err != nil {
		return time, err
	}

	return time, nil
}

func(r *mysqlTimeRepository) Create(time entity.Time) error {
	
	err := r.DB.Create(&time).Error

	if err != nil {
		return err
	}

	return nil
}

func(r *mysqlTimeRepository) Update(time entity.Time) error {
	
	err := r.DB.Save(&time).Error

	if err != nil {
		return err
	}

	return nil
}

func(r *mysqlTimeRepository) Delete(id string) error {
	
	err := r.DB.Delete(&entity.Time{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
