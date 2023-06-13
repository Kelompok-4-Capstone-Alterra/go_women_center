package repository

import (
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(topic, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetAllByPopular(topic, popular, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetAllByCreated(topic, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetById(id string) (*response.ResponseForum, error)
	Delete(id string) error
}

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) GetAll(topic, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if myforum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND forums.topic LIKE ? AND forums.deleted_at IS NULL", myforum, categories, "%"+topic+"%").
		Group("forums.id").Count(&totalData).Offset(offset).Limit(limit).Preload("UserForums").
		Find(&response).Error

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByPopular(topic, popular, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if myforum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND topic LIKE ? AND forums.deleted_at IS NULL", myforum, categories, "%"+topic+"%").
		Group("forums.id").Count(&totalData).
		Order("member " + popular).Offset(offset).Limit(limit).Preload("UserForums").
		Find(&response).Error

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByCreated(topic, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if myforum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND category "+logicOperationCategory+" ? AND topic LIKE ? AND forums.deleted_at IS NULL", myforum, categories, "%"+topic+"%").
		Group("forums.id").
		Order("forums.created_at " + created).Count(&totalData).Offset(offset).Limit(limit).Preload("UserForums").
		Find(&response).Error

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetById(id string) (*response.ResponseForum, error) {
	var forumDetail response.ResponseForum

	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Group("forums.id").Having("forums.id =?", id).Preload("UserForums").
		Find(&forumDetail).Error

	if err != nil {
		return nil, err
	}

	return &forumDetail, nil
}

func (fr mysqlForumRepository) Delete(id string) error {
	err := fr.DB.Where("id = ? ", id).Delete(&entity.Forum{}).Error
	if err != nil {
		return err
	}
	return nil
}
