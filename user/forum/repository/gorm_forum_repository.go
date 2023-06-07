package repository

import (
	"fmt"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(id_user, topic, categories, myforum string) ([]response.ResponseForum, error)
	GetAllByPopular(id_user, topic, popular, categories, myforum string) ([]response.ResponseForum, error)
	GetAllByCreated(id_user, topic, created, categories, myforum string) ([]response.ResponseForum, error)
	GetById(id, user_id string) (*response.ResponseForum, error)
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

func (fr mysqlForumRepository) GetAll(id, topic, categories, myforum string) ([]response.ResponseForum, error) {
	var logicOperationCategory string
	var logicOperationUser string

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
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("category_id "+logicOperationCategory+" ? AND topic LIKE ?", categories, "%"+topic+"%").
		Group("forums.id").Having("forums.user_id "+logicOperationUser+" ?", myforum).
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
		return nil, err
	}
	return response, nil
}

func (fr mysqlForumRepository) GetAllByPopular(id_user, topic, popular, categories, myforum string) ([]response.ResponseForum, error) {
	var logicOperationCategory string
	var logicOperationUser string

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
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("category_id "+logicOperationCategory+" ? AND topic LIKE ?", categories, "%"+topic+"%").
		Group("forums.id").Having("forums.user_id "+logicOperationUser+" ?", myforum).
		Order("member " + popular).
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
		return nil, err
	}
	return response, nil
}

func (fr mysqlForumRepository) GetAllByCreated(id_user, topic, created, categories, myforum string) ([]response.ResponseForum, error) {
	var logicOperationCategory string
	var logicOperationUser string

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
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("category_id "+logicOperationCategory+" ? AND topic LIKE ?", categories, "%"+topic+"%").
		Group("forums.id").Having("forums.user_id "+logicOperationUser+" ?", myforum).
		Order("forums.created_at " + created).
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
		return nil, err
	}
	return response, nil
}

func (fr mysqlForumRepository) GetById(id, user_id string) (*response.ResponseForum, error) {
	var forumDetail response.ResponseForum

	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").
		Group("forums.id").Having("forums.id =?", id).
		Find(&forumDetail).Error

	for i := 0; i < len(forumDetail.UserForums); i++ {
		if forumDetail.UserForums[i].UserId == user_id {
			forumDetail.Status = true
			break
		}
	}
	forumDetail.UserForums = nil

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
	fmt.Println(forum)
	return nil
}

func (fr mysqlForumRepository) Delete(id string) error {
	err := fr.DB.Where("id = ?", id).Take(&entity.Forum{}).Error

	if err != nil {
		return err
	}

	err2 := fr.DB.Delete(&entity.Forum{}, &id).Error
	if err != nil {
		return err2
	}
	return nil
}