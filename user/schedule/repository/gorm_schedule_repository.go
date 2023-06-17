package repository

import (
	"log"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	GetTransactionByCounselorId(counselorId string) ([]entity.Transaction, error)
	GetCurDateByCounselorId(counselorId string) (entity.Date, error)
	GetTimeByCounselorId(counselorId string) ([]entity.Time, error)
	GetDateById(id string) (entity.Date, error)
	GetTimeById(id string) (entity.Time, error)
}

type mysqlScheduleRepository struct {
	DB *gorm.DB
}

func NewMysqlScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &mysqlScheduleRepository{db}
}

func(r *mysqlScheduleRepository) GetTransactionByCounselorId(counselorId string) ([]entity.Transaction, error) {
	
	
	currentTime := time.Now()
	currentDate := currentTime.Format(time.DateOnly)
	
	var transactions []entity.Transaction
	
	err := r.DB.Model(&entity.Transaction{}).Where("DATE(created_at) = ?", currentDate).Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	log.Println(transactions)
	return transactions, nil
}

func(r *mysqlScheduleRepository) GetTimeByCounselorId(counselorId string) ([]entity.Time, error){

	var times []entity.Time

	err := r.DB.Model(&entity.Time{}).Find(&times, "counselor_id = ?", counselorId).Error

	if err != nil {
		return times, err
	}
	
	return times, nil
}

func(r *mysqlScheduleRepository) GetCurDateByCounselorId(counselorId string) (entity.Date, error) {

	var dates entity.Date

	currentTime := time.Now()
	currentDate := currentTime.Format(time.DateOnly)

	err := r.DB.Model(&entity.Date{}).First(&dates, "counselor_id = ? AND date = ?", counselorId, currentDate).Error

	if err != nil {
		return dates, err
	}
	
	return dates, nil

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