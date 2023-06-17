package repository

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(search, topic, sortBy string) ([]counselor.GetAllResponse, error)
	GetById(id string) (counselor.GetByResponse, error)
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func(r *mysqlCounselorRepository) GetAll(search, topic, sortBy string) ([]counselor.GetAllResponse, error) {

	var counselors []counselor.GetAllResponse
	var totalData int64

	currentTime := time.Now()
	currentDate := currentTime.Format(time.DateOnly)

	// get counselor that have date/schedule today
	err := r.DB.Table("counselors").
		Where("topic = ? AND name LIKE ?", topic, "%"+search+"%").
		Joins("INNER JOIN dates ON dates.counselor_id = counselors.id").
		Where("dates.date = ?", currentDate).
		Count(&totalData).
		Order(sortBy).
		Find(&counselors).Error

	if err != nil {
		return nil, err
	}

	return counselors, nil
}

func(r *mysqlCounselorRepository) GetById(id string) (counselor.GetByResponse, error) {
	
	var counselor counselor.GetByResponse

	err := r.DB.Model(&entity.Counselor{}).Where("id = ?", id).First(&counselor).Error
	
	if err != nil {
		return counselor, err
	}

	return counselor, nil
}