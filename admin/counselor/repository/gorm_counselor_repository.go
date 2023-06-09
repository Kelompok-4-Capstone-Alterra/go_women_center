package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(search string, offset, limit int) ([]counselor.GetAllResponse, int64, error)
	GetById(id string) (counselor.GetByResponse, error)
	GetByEmail(email string) (counselor.GetByResponse, error)
	GetByUsername(username string) (counselor.GetByResponse, error)
	Create(counselor entity.Counselor) error
	Update(id string, counselor entity.Counselor) error
	Delete(id string) error
}

type mysqlCounselorRepository struct {
	DB *gorm.DB
}

func NewMysqlCounselorRepository(db *gorm.DB) CounselorRepository{
	return &mysqlCounselorRepository{DB: db}
}

func(r *mysqlCounselorRepository) GetAll(search string, offset, limit int) ([]counselor.GetAllResponse, int64, error) {
	var counselor []counselor.GetAllResponse

	var totalData int64

	err := r.DB.Model(&entity.Counselor{}).
		Where("name LIKE ? OR topic LIKE ? OR username LIKE ? OR email LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalData).
		Offset(offset).Limit(limit).Find(&counselor).Error

	if err != nil {
		return nil, totalData,err
	}
	
	return counselor, totalData, nil
}

func(r *mysqlCounselorRepository) GetById(id string) (counselor.GetByResponse, error) {
	var counselor counselor.GetByResponse
	err := r.DB.Model(&entity.Counselor{}).First(&counselor, "id = ?", id).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func(r *mysqlCounselorRepository) GetByEmail(email string) (counselor.GetByResponse, error) {
	var counselor counselor.GetByResponse
	err := r.DB.Model(&entity.Counselor{}).First(&counselor, "email = ?", email).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func(r *mysqlCounselorRepository) GetByUsername(username string) (counselor.GetByResponse, error) {
	var counselor counselor.GetByResponse
	err := r.DB.Model(&entity.Counselor{}).First(&counselor, "username = ?", username).Error
	if err != nil {
		return counselor, err
	}
	return counselor, nil
}

func(r *mysqlCounselorRepository) Create(counselor entity.Counselor) error {
	err := r.DB.Create(&counselor).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *mysqlCounselorRepository) Update(id string, counselor entity.Counselor) error {
	
	err := r.DB.Model(&entity.Counselor{}).Where("id = ?", id).Updates(counselor).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *mysqlCounselorRepository) Delete(id string) error {

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Review{}).Where("counselor_id = ?", id).Unscoped().Delete(&entity.Review{}).Error
		
		if err != nil {
			return err
		}

		err = tx.Model(&entity.Date{}).Where("counselor_id = ?", id).Delete(&entity.Date{}).Error

		if err != nil {
			return err
		}

		err = tx.Model(&entity.Time{}).Where("counselor_id = ?", id).Delete(&entity.Time{}).Error

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
