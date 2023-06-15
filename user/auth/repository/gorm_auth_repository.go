package repository

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	userError "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userData entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
	GetByUsernameAndEmail(username, email string) (entity.User, error)
	GetById(id string) (entity.User, error)
}

type userGormMysqlRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
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

func (u *userGormMysqlRepo) GetByUsername(username string) (entity.User, error) {
	savedUser := entity.User{}
	err := u.DB.Where("username = ?", username).First(&savedUser).Error
	if err != nil {
		return entity.User{}, err
	}

	return savedUser, nil
}

func (u *userGormMysqlRepo) GetById(id string) (entity.User, error) {
	savedUser := entity.User{}
	err := u.DB.Where("id = ?", id).First(&savedUser).Error
	if err != nil {
		return entity.User{}, err
	}

	return savedUser, nil
}

func (u *userGormMysqlRepo) GetByUsernameAndEmail(username, email string) (entity.User, error) {
	savedUser := entity.User{}
	err := u.DB.Where("username = ? OR email = ?", username, email).First(&savedUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, userError.ErrDataNotFound
	}
	
	if err != nil {
		return entity.User{}, err
	}

	return savedUser, nil
}