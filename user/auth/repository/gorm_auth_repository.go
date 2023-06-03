package repository

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	userError "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth"
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.User{}, userError.ErrUserIsRegistered
		}
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
