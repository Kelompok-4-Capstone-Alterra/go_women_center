package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	DeleteByCounselorId(counselorId string) error
	Create(dates []entity.Date, times []entity.Time) error
	GetByCounselorId(counselorId string) (schedule.GetAllResponse, error)
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

func (r *mysqlScheduleRepository) GetByCounselorId(counselorId string) (schedule.GetAllResponse, error) {

	var dates []entity.Date
	var times []entity.Time

	var schedule schedule.GetAllResponse

	err := r.DB.Find(&dates, "counselor_id = ?", counselorId).Error

	if err != nil {
		return schedule, err
	}

	err = r.DB.Find(&times, "counselor_id = ?", counselorId).Error

	if err != nil {
		return schedule, err
	}

	var datesRes = make([]string, len(dates))
	var timesRes = make([]string, len(times))

	for i, date := range dates {
		datesRes[i] = date.Date.Format("2006-01-02")
	}

	for i, time := range times {
		timesRes[i] = time.Time
	}

	schedule.Dates = datesRes
	schedule.Times = timesRes
	
	return schedule, nil
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
		return err
	}

	return nil
}