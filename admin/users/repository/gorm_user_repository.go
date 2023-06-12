package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id string) (users.GetByIdResponse, error)
	GetAll(search string, offset, limit int) ([]users.GetAllResponse, int64, error)
	Delete(id string) error
}

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) UserRepository {
	return &mysqlUserRepository{DB: db}
}

func (r *mysqlUserRepository) GetById(id string) (users.GetByIdResponse, error) {
	var userRes users.GetByIdResponse
	err := r.DB.Model(&entity.User{}).Where("id = ?", id).First(&userRes).Error
	return userRes, err
}

func (r *mysqlUserRepository) GetAll(search string, offset, limit int) ([]users.GetAllResponse, int64, error) {
	var usersRes []users.GetAllResponse
	var totalData int64
	err := r.DB.Model(&entity.User{}).
		Select("id, name, email, username").
		Where("name LIKE ? OR email LIKE ? OR username LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&totalData).
		Offset(offset).
		Limit(limit).
		Find(&usersRes).Error

	if err != nil {
		return []users.GetAllResponse{}, 0, err
	}

	return usersRes, totalData ,nil
}

func (r *mysqlUserRepository) Delete(id string) error {
	
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Review{}).Unscoped().Delete(&entity.Review{}, "user_id = ?", id).Error

		if err != nil {
			return err
		}
		err = tx.Model(&entity.User{}).Unscoped().Delete(&entity.User{}, "id = ?", id).Error

		if err != nil {
			return err
		}
		
		return nil
	})

	return err
}