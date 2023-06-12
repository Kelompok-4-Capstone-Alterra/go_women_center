package repository

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(id_user, topic, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetAllByPopular(id_user, topic, popular, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetAllByCreated(id_user, topic, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error)
	GetById(id, user_id string) (*response.ResponseForum, error)
	Create(forum *entity.Forum) error
	Update(id, user_id string, forumId *entity.Forum) error
	Delete(id, user_id string) error
}

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) GetAll(id, topic, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
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

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == id {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, errors.New("failed to get all forum data")
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByPopular(id_user, topic, popular, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
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

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == id_user {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, errors.New("failed to get all forum data")
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByCreated(id_user, topic, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int64, error) {
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

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == id_user {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, errors.New("failed to get all forum data")
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetById(id, user_id string) (*response.ResponseForum, error) {
	var forumDetail response.ResponseForum

	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Group("forums.id").Having("forums.id =?", id).Preload("UserForums").
		Find(&forumDetail).Error

	for i := 0; i < len(forumDetail.UserForums); i++ {
		if forumDetail.UserForums[i].UserId == user_id {
			forumDetail.Status = true
			break
		}
	}
	forumDetail.UserForums = nil

	if forumDetail.ID == "" {
		return nil, errors.New("invalid id user " + id)
	} else if err != nil {
		return nil, errors.New("failed to get forum data details")
	}

	return &forumDetail, nil
}

func (fr mysqlForumRepository) Create(forum *entity.Forum) error {
	err := fr.DB.Save(forum).Error

	if err != nil {
		return errors.New("failed created forum data")
	}
	return nil
}

func (fr mysqlForumRepository) Update(id, user_id string, forumId *entity.Forum) error {
	var forum entity.Forum
	err2 := fr.DB.Model(&forum).Where("id = ? AND user_id = ? ", id, user_id).Updates(&forumId).Error

	if err2 != nil {
		return errors.New("failed to updated forum data")
	}

	return nil
}

func (fr mysqlForumRepository) Delete(id, user_id string) error {
	err2 := fr.DB.Where("id = ? AND user_id = ? ", id, user_id).Delete(&entity.Forum{}).Error
	if err2 != nil {
		return errors.New("failed to delete forum data")
	}
	return nil
}
