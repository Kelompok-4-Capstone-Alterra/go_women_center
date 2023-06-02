package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(userData entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
}

type userGormMysqlRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userGormMysqlRepo {
	return &userGormMysqlRepo{
		DB: db,
	}
}

func (u *userGormMysqlRepo) Create(userData entity.User) (entity.User, error) {
	err := u.DB.Debug().Create(&userData).Error
	if err != nil {
		return entity.User{}, err
	}
	return userData, nil
}

func (u *userGormMysqlRepo) GetByEmail(email string) (entity.User, error) {
	savedUser := entity.User{}
	err := u.DB.Where("email = ?", email).First(&savedUser).Error
	if err != nil {
		return entity.User{}, err
	}

	return savedUser, nil
}
