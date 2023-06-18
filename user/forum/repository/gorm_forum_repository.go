package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"gorm.io/gorm"
)

type ForumRepository interface {
	GetAll(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetAllSortBy(getAllRequest forum.GetAllRequest, category string) ([]forum.ResponseForum, int64, error)
	GetAllByPopular(id_user, topic, popular, categories, myforum string, offset, limit int) ([]forum.ResponseForum, int64, error)
	GetAllByCreated(id_user, topic, created, categories, myforum string, offset, limit int) ([]forum.ResponseForum, int64, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
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

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == getAllRequest.UserId {
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

func (fr mysqlForumRepository) GetAllByPopular(id_user, topic, popular, categories, myforum string, offset, limit int) ([]forum.ResponseForum, int64, error) {
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

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.user_id "+logicOperationUser+" ? AND forums.category "+logicOperationCategory+" ? AND topic LIKE ?", myforum, categories, "%"+topic+"%").
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

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(response[i].UserForums); j++ {
			if response[i].UserForums[j].UserId == getAllRequest.UserId {
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

func (fr mysqlForumRepository) GetAllByCreated(id_user, topic, created, categories, myforum string, offset, limit int) ([]forum.ResponseForum, int64, error) {
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

	var response []forum.ResponseForum
	err := fr.DB.Table("forums").
		Select("forums.id, forums.user_id, forums.category, forums.link, forums.topic, COUNT(user_forums.id) AS member, forums.created_at").
		Joins("LEFT JOIN user_forums ON forums.id = user_forums.forum_id").Where("forums.user_id "+logicOperationUser+" ? AND category "+logicOperationCategory+" ? AND topic LIKE ?", myforum, categories, "%"+topic+"%").
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

func (fr mysqlForumRepository) Update(id, user_id string, forumId *entity.Forum) error {
	var forum entity.Forum
	err := fr.DB.Model(&forum).Where("id = ? AND user_id = ? ", id, user_id).Updates(&forumId).Error
	if err != nil {
		return err
	}

	return nil
}

func (fr mysqlForumRepository) Delete(id, user_id string) error {
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
