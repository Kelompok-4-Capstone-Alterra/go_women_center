package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type AdminRepo interface {
	GetByUsername(username string) (entity.Admin, error)
}

type adminGormMysqlRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *adminGormMysqlRepo {
	return &adminGormMysqlRepo{
		DB: db,
	}
}

func (a *adminGormMysqlRepo) GetByUsername(username string) (entity.Admin, error) {
	adminData := entity.Admin{}
	err := a.DB.Where("username = ?", username).First(&adminData).Error
	if err != nil {
		return entity.Admin{}, err
	}

	return adminData, nil
}