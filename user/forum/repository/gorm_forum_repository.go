package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error)
	GetAllByPopular(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error)
	GetAllByCreated(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
	Create(createForum *entity.Forum) error
	Update(id, user_id string, forumId *forum.UpdateRequest) error
	Delete(id, user_id string) error
}

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) GetAll(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if getAllParam.MyForum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND forums.topic LIKE ? AND forums.deleted_at IS NULL", getAllParam.MyForum, categories, "%"+getAllParam.Topic+"%").
		Group("forums.id").Count(&totalData).Offset(getAllParam.Offset).Limit(getAllParam.Limit).Preload("UserForums").
		Find(&response).Error

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == getAllParam.IdUser {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByPopular(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if getAllParam.MyForum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND topic LIKE ? AND forums.deleted_at IS NULL", getAllParam.MyForum, categories, "%"+getAllParam.Topic+"%").
		Group("forums.id").Count(&totalData).
		Order("member " + getAllParam.Popular).Offset(getAllParam.Offset).Limit(getAllParam.Limit).Preload("UserForums").
		Find(&response).Error

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == getAllParam.IdUser {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetAllByCreated(getAllParam forum.GetAllRequest, categories string) ([]forum.ResponseForum, int64, error) {
	var logicOperationCategory string
	var logicOperationUser string
	var totalData int64

	if categories == "" {
		logicOperationCategory = "!="
	} else {
		logicOperationCategory = "="
	}

	if getAllParam.MyForum == "" {
		logicOperationUser = "!="
	} else {
		logicOperationUser = "="
	}

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, users.name,users.profile_picture, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at, forums.updated_at,forums.deleted_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").
		Joins("LEFT JOIN users ON forums.user_id = users.id").
		Where("forums.user_id "+logicOperationUser+" ? AND category "+logicOperationCategory+" ? AND topic LIKE ? AND forums.deleted_at IS NULL", getAllParam.MyForum, categories, "%"+getAllParam.Topic+"%").
		Group("forums.id").
		Order("forums.created_at " + getAllParam.Created).Count(&totalData).Offset(getAllParam.Offset).Limit(getAllParam.Limit).Preload("UserForums").
		Find(&response).Error

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == getAllParam.IdUser {
				response[i].Status = true
				break
			}
		}
		response[i].UserForums = nil
	}

	if err != nil {
		return nil, totalData, err
	}

	return response, totalData, nil
}

func (fr mysqlForumRepository) GetById(id, user_id string) (*forum.ResponseForum, error) {
	var forumDetail forum.ResponseForum

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

	if err != nil {
		return nil, err
	}

	return &forumDetail, nil
}

func (fr mysqlForumRepository) Create(createForum *entity.Forum) error {
	err := fr.DB.Save(createForum).Error

	if err != nil {
		return err
	}
	return nil
}

func (fr mysqlForumRepository) Update(id, user_id string, forumId *forum.UpdateRequest) error {
	var forum entity.Forum
	err := fr.DB.Model(&forum).Where("id = ? AND user_id = ? ", id, user_id).Updates(&forumId).Error

	if err != nil {
		return err
	}

	return nil
}

func (fr mysqlForumRepository) Delete(id, user_id string) error {
	err := fr.DB.Where("id = ? AND user_id = ? ", id, user_id).Delete(&entity.Forum{}).Error
	if err != nil {
		return err
	}
	return nil
}
