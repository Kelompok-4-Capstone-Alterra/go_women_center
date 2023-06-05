package repository

import (
	"fmt"
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(topic string) ([]response.ResponseForum, error)
	GetAllByPopular(topic, popular string) ([]response.ResponseForum, error)
	GetAllByCreated(topic, created string) ([]response.ResponseForum, error)
	GetByCategory(category_id string) ([]response.ResponseForum, error)
	GetByMyForum(id_user string) ([]response.ResponseForum, error)
	GetById(id string) (*response.ResponseForumDetail, error)
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

func (fr mysqlForumRepository) GetAll(topic string) ([]response.ResponseForum, error) {
	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("topic LIKE ?", "%"+topic+"%").
		Group("forums.id").
		Find(&response).Error

	idUserJWT := 1
	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == uint(idUserJWT) {
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

func (fr mysqlForumRepository) GetAllByPopular(topic, popular string) ([]response.ResponseForum, error) {
	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("topic LIKE ?", "%"+topic+"%").
		Group("forums.id").Order("member " + popular).
		Find(&response).Error

	idUserJWT := 1
	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == uint(idUserJWT) {
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

func (fr mysqlForumRepository) GetAllByCreated(topic, created string) ([]response.ResponseForum, error) {
	fmt.Println("GetAllByCreated :", created)
	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("topic LIKE ?", "%"+topic+"%").
		Group("forums.id").Order("forums.created_at " + created).
		Find(&response).Error

	idUserJWT := 1
	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == uint(idUserJWT) {
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

func (fr mysqlForumRepository) GetByCategory(category_id string) ([]response.ResponseForum, error) {
	var response []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").Where("category_id = ?", category_id).
		Group("forums.id").
		Find(&response).Error

	idUserJWT := 1
	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == uint(idUserJWT) {
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

func (fr mysqlForumRepository) GetByMyForum(id_user string) ([]response.ResponseForum, error) {
	var forum []response.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category_id, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Preload("UserForums").
		Group("forums.id").
		Find(&forum).Error

	id, _ := strconv.Atoi(id_user)
	var myForum []response.ResponseForum
	for i := 0; i < len(forum); i++ {
		for j := 0; j < len(forum[i].UserForums); j++ {
			if forum[i].UserForums[j].UserId == uint(id) {
				forum[i].Status = true
				myForum = append(myForum, forum[i])
				myForum[len(myForum)-1].UserForums = nil
				break
			}
		}
	}

	if err != nil {
		return nil, err
	}
	return myForum, nil
}

func (fr mysqlForumRepository) GetById(id string) (*response.ResponseForumDetail, error) {
	var forumdetail response.ResponseForumDetail
	err := fr.DB.Model(entity.Forum{}).Preload("UserForums").First(&forumdetail, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	forumdetail.UserForums = nil
	return &forumdetail, nil
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
