package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(offset, limit int, topic string) ([]counselor.GetAllResponse, error)
	Count() (int, error)
	GetById(id string) (counselor.GetByResponse, error)
	Search(search, topic string, offset, limit int) ([]counselor.GetAllResponse, error)
	CountBySearch(search, topic string) (int, error)
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func(r *mysqlCounselorRepository) GetAll(offset, limit int, topic string) ([]counselor.GetAllResponse, error) {

	var counselors []counselor.GetAllResponse

	err := r.DB.Model(&entity.Counselor{}).Where("topic = ?", topic).Offset(offset).Limit(limit).Find(&counselors).Error

	if err != nil {
		return nil, err
	}

	return counselors, nil
}

func(r *mysqlCounselorRepository) Count() (int, error) {

	var totalData int64

	err := r.DB.Model(&entity.Counselor{}).Count(&totalData).Error

	if err != nil {
		return 0, err
	}

	return int(totalData), nil
}

func(r *mysqlCounselorRepository) GetById(id string) (counselor.GetByResponse, error) {
	
	var counselor counselor.GetByResponse

	err := r.DB.Model(&entity.Counselor{}).Where("id = ?", id).First(&counselor).Error
	
	if err != nil {
		return counselor, err
	}

	return counselor, nil
}

func(r *mysqlCounselorRepository) CountBySearch(search, topic string) (int, error) {
	
	var totalData int64

	err := r.DB.Model(&entity.Counselor{}).Where("name LIKE ? AND topic = ?", "%"+search+"%", topic).Count(&totalData).Error

	if err != nil {
		return 0, err
	}

	return int(totalData), nil
}

func(r *mysqlCounselorRepository) Search(search, topic string, offset, limit int) ([]counselor.GetAllResponse, error) {
	
	var counselors []counselor.GetAllResponse

	err := r.DB.Model(&entity.Counselor{}).Where("name LIKE ? AND topic = ?", "%"+search+"%", topic).Offset(offset).Limit(limit).Find(&counselors).Error

	if err != nil {
		return nil, err
	}

	return counselors, nil
}