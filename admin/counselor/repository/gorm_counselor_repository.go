package repository

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CounselorRepository interface {
	GetAll(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error)
	GetAllHasSchedule(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error)
	GetAllNotHasSchedule(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error)
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

func(r *mysqlCounselorRepository) GetAll(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error) {
	var counselor []counselor.GetAllResponse

	var totalData int64

	err := r.DB.Model(&entity.Counselor{}).
		Where("name LIKE ? OR topic LIKE ? OR username LIKE ? OR email LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalData).
		Order(sortBy).
		Offset(offset).Limit(limit).Find(&counselor).Error

	if err != nil {
		return nil, totalData,err
	}
	
	return counselor, totalData, nil
}

func(r *mysqlCounselorRepository) GetAllHasSchedule(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error) {

	var counselor []counselor.GetAllResponse

	var totalData int64
	// get counselor that has schedule
	err := r.DB.Table("counselors").
		Where("deleted_at IS NULL").
		Joins("INNER JOIN dates ON counselors.id = dates.counselor_id").
		Where("counselors.name LIKE ? OR counselors.topic LIKE ? OR counselors.username LIKE ? OR counselors.email LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Group("counselors.id").
		Count(&totalData).
		Order(sortBy).
		Offset(offset).
		Limit(limit).
		Find(&counselor).Error

	if err != nil {
		log.Println(err)
		return nil, totalData, err
	}
	
	return counselor, totalData, nil
}

func(r *mysqlCounselorRepository) GetAllNotHasSchedule(search, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int64, error) {

	var counselor []counselor.GetAllResponse

	var totalData int64

	// get counselor that not has schedule
	err := r.DB.Table("counselors").
		Where("deleted_at IS NULL").
		Joins("LEFT JOIN dates ON counselors.id = dates.counselor_id").
		Where("dates.counselor_id IS NULL").
		Where("counselors.name LIKE ? OR counselors.topic LIKE ? OR counselors.username LIKE ? OR counselors.email LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Group("counselors.id").
		Count(&totalData).
		Order(sortBy).
		Offset(offset).
		Limit(limit).
		Find(&counselor).Error

	if err != nil {
		log.Println(err)
		return nil, totalData, err
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
		err := tx.Model(&entity.Review{}).Where("counselor_id = ?", id).Delete(&entity.Review{}).Error
		
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

		err = tx.Model(&entity.Counselor{}).Where("id = ?", id).Delete(&entity.Counselor{}).Error

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