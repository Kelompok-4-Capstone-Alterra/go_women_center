package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
)

type ForumUsecaseInterface interface {
	GetAll(getAllRequest forum.GetAllRequest) ([]forum.ResponseForum, int, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
	Create(createRequest *forum.CreateRequest) error
	Update(id, user_id string, updateRequest *forum.UpdateRequest) error
	Delete(id, user_id string) error
}

type ForumUsecase struct {
	ForumR repository.ForumRepository
}

func NewForumUsecase(ForumR repository.ForumRepository) ForumUsecaseInterface {
	return &ForumUsecase{
		ForumR: ForumR,
	}
}

func (fu ForumUsecase) GetAll(getAllRequest forum.GetAllRequest) ([]forum.ResponseForum, int, error) {
	var forums []forum.ResponseForum
	var err error
	var totalData int64

	switch getAllRequest.SortBy {
	case "oldest":
		getAllRequest.SortBy = "forums.created_at ASC"
	case "newest":
		getAllRequest.SortBy = "forums.created_at DESC"
	case "popular":
		getAllRequest.SortBy = "member DESC"
	}

	switch getAllRequest.MyForum {
	case "yes":
		getAllRequest.MyForum = getAllRequest.UserId
	}

	var newCategory string
	category, ok := constant.TOPICS[getAllRequest.CategoryId]
	if ok {
		newCategory = category[0]
	}

	if getAllRequest.SortBy != "" {
		forums, totalData, err = fu.ForumR.GetAllSortBy(getAllRequest, newCategory)
	} else {
		forums, totalData, err = fu.ForumR.GetAll(getAllRequest, newCategory)
	}

	if err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllRequest.Limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*forum.ResponseForum, error) {
	forumId, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, forum.ErrFailedGetDetailReadingList
	} else if forumId.ID == "" {
		return nil, forum.ErrInvalidId
	}
	return forumId, nil
}

func (fu ForumUsecase) Create(createRequest *forum.CreateRequest) error {
	var newCategory string
	category, ok := constant.TOPICS[createRequest.CategoryId]
	if ok {
		newCategory = category[0]
	}

	var validUrlHost = map[string]bool{
		"t.me": true,
	}
	err := helper.IsValidUrl(createRequest.Link, validUrlHost)
	if err == helper.ErrInvalidUrl {
		return helper.ErrInvalidUrl
	} else if err == helper.ErrInvalidUrlHost {
		return forum.ErrInvalidUrlHost
	}

	createForum := entity.Forum{
		ID:       createRequest.ID,
		UserId:   createRequest.UserId,
		Category: newCategory,
		Link:     createRequest.Link,
		Topic:    createRequest.Topic,
		Status:   createRequest.Status,
		Member:   createRequest.Member,
	}

	err = fu.ForumR.Create(&createForum)
	if err != nil {
		return forum.ErrFailedCreateReadingList
	}
	return nil
}

func (fu ForumUsecase) Update(id, user_id string, updateRequest *forum.UpdateRequest) error {
	forumId, err := fu.ForumR.GetById(id, user_id)
	if err != nil {
		return err
	} else if forumId.ID == "" {
		return forum.ErrInvalidId
	} else if forumId.UserId != user_id {
		return forum.ErrNotAccess
	}

	var newCategory string
	category, ok := constant.TOPICS[updateRequest.CategoryId]
	if ok {
		newCategory = category[0]
	}

	updateForum := entity.Forum{
		Category: newCategory,
		Link:     updateRequest.Link,
		Topic:    updateRequest.Topic,
	}
	err = fu.ForumR.Update(id, &updateForum)

	if err != nil {
		return forum.ErrFailedUpdateReadingList
	}
	return nil
}

func (fu ForumUsecase) Delete(id, user_id string) error {
	forumId, err := fu.ForumR.GetById(id, user_id)
	if err != nil {
		return err
	} else if forumId.ID == "" {
		return forum.ErrInvalidId
	} else if forumId.UserId != user_id {
		return forum.ErrNotAccess
	}

	err = fu.ForumR.Delete(id)

	if err != nil {
		return forum.ErrFailedDeleteReadingList
	}
	return nil
}
