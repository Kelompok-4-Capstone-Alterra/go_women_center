package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id string) (profile.GetByIdResponse, error)
	// Update(userData entity.User) error
}

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) UserRepository {
	return &mysqlUserRepository{DB}

}

func(u *mysqlUserRepository) GetById(id string) (profile.GetByIdResponse, error) {
	
	var profile profile.GetByIdResponse

	err := u.DB.Model(&entity.User{}).First(&profile,"id = ?", id).Error

	if err != nil {
		return profile, err
	}

	return profile, nil	
}

// func(u *mysqlUserRepository) Update(userData entity.User) error{

// }