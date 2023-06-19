package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	userForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum"
	"gorm.io/gorm"
)

type UserForumRepository interface {
	GetById(user_id, forum_id string) (*userForum.Response, error)
	Create(user_forum *entity.UserForum) error
}

type mysqlUserForumRepository struct {
	DB *gorm.DB
}

func NewMysqlUserForumRepository(db *gorm.DB) UserForumRepository {
	return &mysqlUserForumRepository{DB: db}
}

func (ufr mysqlUserForumRepository) GetById(user_id, forum_id string) (*userForum.Response, error) {
	var response userForum.Response
	err := ufr.DB.Model(&entity.UserForum{}).Where("user_id = ? AND forum_id = ?", user_id, forum_id).First(&response).Error

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (ufr mysqlUserForumRepository) Create(user_forum *entity.UserForum) error {
	err := ufr.DB.Save(user_forum).Error

	if err != nil {
		return err
	}
	return nil
}
