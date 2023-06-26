package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id string) (entity.User, error)
	Update(userData entity.User) error
	GetByEmail(email string) error
	GetByUsername(username string) error
}

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) UserRepository {
	return &mysqlUserRepository{DB}

}

func(u *mysqlUserRepository) GetByEmail(email string) error {
	return u.DB.First(&entity.User{},"email = ?", email).Error
}

func(u *mysqlUserRepository) GetByUsername(username string) error {
	return u.DB.First(&entity.User{},"username = ?", username).Error
}

func(u *mysqlUserRepository) GetById(id string) (entity.User, error) {
	
	var userData entity.User

	err := u.DB.Model(&entity.User{}).First(&userData,"id = ?", id).Error

	if err != nil {
		return userData, err
	}

	return userData, nil	
}

func(u *mysqlUserRepository) Update(userData entity.User) error{

	err := u.DB.Model(&entity.User{}).Where("id = ?", userData.ID).Updates(userData).Error

	if err != nil {
		return err
	}

	return nil
}