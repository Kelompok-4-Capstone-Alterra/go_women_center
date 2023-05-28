package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"gorm.io/gorm"
)

type AdminRepo interface {
	GetByEmail(email string) (domain.Admin, error)
}

type adminGormMysqlRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *adminGormMysqlRepo {
	return &adminGormMysqlRepo{
		DB: db,
	}
}

func (a *adminGormMysqlRepo) GetByEmail(email string) (domain.Admin, error) {
	adminData := domain.Admin{}
	err := a.DB.Where("email = ?", email).First(&adminData).Error
	if err != nil {
		return domain.Admin{}, err
	}

	return adminData, nil
}