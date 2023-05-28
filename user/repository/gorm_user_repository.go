package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(userData domain.User) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
}

type userGormMysqlRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userGormMysqlRepo {
	return &userGormMysqlRepo{
		DB: db,
	}
}

func (u *userGormMysqlRepo) Create(userData domain.User) (domain.User, error) {
	err := u.DB.Debug().Create(&userData).Error
	if err != nil {
		return domain.User{}, err
	}
	return userData, nil
}

func (u *userGormMysqlRepo) GetByEmail(email string) (domain.User, error) {
	savedUser := domain.User{}
	err := u.DB.Where("email = ?", email).First(&savedUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return savedUser, nil
}
