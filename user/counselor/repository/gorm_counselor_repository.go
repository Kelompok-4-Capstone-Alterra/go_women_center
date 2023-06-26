package repository

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(search, topic, sortBy, is_available string) ([]counselor.GetAllResponse, error)
	GetById(id string) (counselor.GetByResponse, error)
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func(r *mysqlCounselorRepository) GetAll(search, topic, sortBy, is_available string) ([]counselor.GetAllResponse, error) {

	var counselors []counselor.GetAllResponse
	
	dbQuery := r.DB.Table("counselors").
		Where("name LIKE ? OR CAST(price AS CHAR) LIKE ? OR CAST(rating AS CHAR) LIKE ?", 
		"%"+search+"%", "%"+search+"%", "%"+search+"%")

	if topic != "" {
		dbQuery.Where("topic = ?", topic)
	}
	
	if is_available == "true" {
		currentTime := time.Now()
		currentDate := currentTime.Format(time.DateOnly)

		dbQuery.Joins("INNER JOIN dates ON dates.counselor_id = counselors.id").
		Where("dates.date = ?", currentDate)
	}
	
	err := dbQuery.Order(sortBy).Find(&counselors).Error

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