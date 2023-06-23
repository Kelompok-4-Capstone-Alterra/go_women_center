package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetAllSortBy(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
	Create(forum *entity.Forum) error
	Update(id string, forumId *entity.Forum) error
	Delete(id string) error
}

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) GetAll(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if category == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if getAllRequest.MyForum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND forums.topic LIKE ?", getAllRequest.MyForum, category, "%"+getAllRequest.Topic+"%").
		Group("forums.id").Count(&totalData).Offset(getAllRequest.Offset).Limit(getAllRequest.Limit).Preload("UserForums").
		Find(&response).Error

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllSortBy(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if category == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if getAllRequest.MyForum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND topic LIKE ?", getAllRequest.MyForum, category, "%"+getAllRequest.Topic+"%").
		Group("forums.id").Count(&totalData).
		Order(getAllRequest.SortBy).Offset(getAllRequest.Offset).Limit(getAllRequest.Limit).Preload("UserForums").
		Find(&response).Error

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetById(id, user_id string) (*forum.ResponseForum, error) {
	var forumDetail forum.ResponseForum

	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").
		Group("forums.id").Having("forums.id =?", id).
		Find(&forumDetail).Error

	if err != nil {
		return nil, err
	}

	return &forumDetail, nil
}

func (fr mysqlForumRepository) Create(forum *entity.Forum) error {
	err := fr.DB.Save(forum).Error

	if err != nil {
		return err
	}
	return nil
}

func (fr mysqlForumRepository) Update(id string, forumId *entity.Forum) error {
	var forum entity.Forum
	err := fr.DB.Model(&forum).Where("id = ?", id).Updates(&forumId).Error
	if err != nil {
		return err
	}

	return nil
}

func (fr mysqlForumRepository) Delete(id string) error {
	err := fr.DB.Transaction(func(tx *gorm.DB) error {
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
