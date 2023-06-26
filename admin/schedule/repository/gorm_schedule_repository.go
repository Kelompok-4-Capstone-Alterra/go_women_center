package repository

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	DeleteByCounselorId(counselorId string) error
	Create(dates []entity.Date, times []entity.Time) error
	GetTimeByCounselorId(counselorId string) ([]entity.Time, error)
	GetDateByCounselorId(counselorId string) ([]entity.Date, error)
	Update(counselorId string, dates []entity.Date, times []entity.Time) error
}

type mysqlScheduleRepository struct {
	DB *gorm.DB
}

func NewMysqlScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &mysqlScheduleRepository{db}
}

func (r *mysqlScheduleRepository) DeleteByCounselorId(counselorId string) error {
	
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&entity.Date{}, "counselor_id = ?", counselorId).Error

		if err != nil {
			return err
		}

		err = tx.Delete(&entity.Time{}, "counselor_id = ?", counselorId).Error
		if err != nil {
			return err
		}

		return nil
	})
	

	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlScheduleRepository) Create(dates []entity.Date, times []entity.Time) error {
	
	err := r.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Model(&entity.Date{}).Create(&dates).Error

		if err != nil {
			return err
		}

		err = tx.Model(&entity.Time{}).Create(&times).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlScheduleRepository) GetTimeByCounselorId(counselorId string) ([]entity.Time, error) {
	var times []entity.Time
	err := r.DB.Find(&times, "counselor_id = ?", counselorId).Error

	if err != nil {
		return times, err
	}
	return times, nil
}

func (r *mysqlScheduleRepository) GetDateByCounselorId(counselorId string) ([]entity.Date, error) {
	var dates []entity.Date
	err := r.DB.Find(&dates, "counselor_id = ?", counselorId).Error
	if err != nil {
		return dates, err
	}
	return dates, nil
}

func (r *mysqlScheduleRepository) Update(counselorId string, dates []entity.Date, times []entity.Time) error {
	
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&entity.Date{}, "counselor_id = ?", counselorId).Error

		if err != nil {
			return err
		}

		err = tx.Delete(&entity.Time{}, "counselor_id = ?", counselorId).Error
		if err != nil {
			return err
		}
		
		err = tx.Create(&dates).Error

		if err != nil {
			return err
		}

		err = tx.Create(&times).Error

		if err != nil {
			return err
		}

		return nil
	})
	

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}