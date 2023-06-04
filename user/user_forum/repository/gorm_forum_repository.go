package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type UserForumRepository interface {
	Create(user_forum *entity.UserForum) error
}

type mysqlUserForumRepository struct {
	DB *gorm.DB
}

func NewMysqlUserForumRepository(db *gorm.DB) UserForumRepository {
	return &mysqlUserForumRepository{DB: db}
}

func (fr mysqlUserForumRepository) Create(user_forum *entity.UserForum) error {
	err := fr.DB.Save(user_forum).Error

	if err != nil {
		return err
	}
	return nil
}
