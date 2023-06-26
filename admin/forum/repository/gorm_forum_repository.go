package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ForumAdminRepository interface {
	GetAll(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetAllSortBy(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetById(id string) (*forum.ResponseForum, error)
	Delete(id string) error
}

type mysqlForumAdminRepository struct {
	DB *gorm.DB
}

func NewMysqlForumAdminRepository(db *gorm.DB) ForumAdminRepository {
	return &mysqlForumAdminRepository{DB: db}
}

func (frar mysqlForumAdminRepository) GetAll(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var totalData int64

	if category == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	var response []forum.ResponseForum
	err := frar.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.category "+logicOperationCategory+" ? AND forums.topic LIKE ?", category, "%"+getAllRequest.Topic+"%").
		Group("forums.id").Count(&totalData).Offset(getAllRequest.Offset).Limit(getAllRequest.Limit).Preload("UserForums").
		Find(&response).Error

	for i := 0; i < len(response); i++ {
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (frar mysqlForumAdminRepository) GetAllSortBy(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var totalData int64

	if category == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	var response []forum.ResponseForum
	err := frar.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.category "+logicOperationCategory+" ? AND topic LIKE ?", category, "%"+getAllRequest.Topic+"%").
		Group("forums.id").Count(&totalData).
		Order(getAllRequest.SortBy).Offset(getAllRequest.Offset).Limit(getAllRequest.Limit).Preload("UserForums").
		Find(&response).Error

	for i := 0; i < len(response); i++ {
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (frar mysqlForumAdminRepository) GetById(id string) (*forum.ResponseForum, error) {
	var forumDetail forum.ResponseForum

	err := frar.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").
		Group("forums.id").Having("forums.id =?", id).
		Find(&forumDetail).Error

	forumDetail.UserForums = nil

	if err != nil {
		return nil, err
	}

	return &forumDetail, nil
}

func (frar mysqlForumAdminRepository) Delete(id string) error {
	err := frar.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.UserForum{}).Unscoped().Delete(&entity.UserForum{}, "forum_id = ?", id).Error

		if err != nil {
			return err
		}
		err = tx.Model(&entity.Forum{}).Unscoped().Delete(&entity.Forum{}, "id = ?", id).Error

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
