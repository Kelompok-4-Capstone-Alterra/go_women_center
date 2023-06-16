package schedule

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	GetDateById(id string) (entity.Date, error)
	GetTimeById(id string) (entity.Time, error)
}

type mysqlScheduleRepository struct {
	DB *gorm.DB
}

func NewMysqlScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &mysqlScheduleRepository{
		DB: db,
	}
}

func (sr *mysqlScheduleRepository) GetDateById(id string) (entity.Date, error) {
	savedDate := entity.Date{}
	err := sr.DB.Where("id = ?", id).First(&savedDate).Error
	if err != nil {
		return entity.Date{}, err
	}

	return savedDate, nil
}

func (sr *mysqlScheduleRepository) GetTimeById(id string) (entity.Time, error) {
	savedTime := entity.Time{}
	err := sr.DB.Where("id = ?", id).First(&savedTime).Error
	if err != nil {
		return entity.Time{}, err
	}

	return savedTime, nil
}