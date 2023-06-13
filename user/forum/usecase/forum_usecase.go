package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
)

type ForumUsecaseInterface interface {
	GetAll(getAllParam forum.GetAllRequest) ([]forum.ResponseForum, int, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
	Create(createForum *forum.CreateRequest) error
	Update(id, user_id string, forumId *forum.UpdateRequest) error
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

func (fu ForumUsecase) GetAll(getAllParam forum.GetAllRequest) ([]forum.ResponseForum, int, error) {
	var forums []forum.ResponseForum
	var err error
	var totalData int64

	categories := constant.TOPICS[getAllParam.Categories]

	if getAllParam.Created == "asc" || getAllParam.Created == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByCreated(getAllParam, categories)
	} else if getAllParam.Popular == "asc" || getAllParam.Popular == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByPopular(getAllParam, categories)
	} else {
		forums, totalData, err = fu.ForumR.GetAll(getAllParam, categories)
	}

	if err != nil {
		return nil, 0, forum.ErrFailedGetForum
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllParam.Limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*forum.ResponseForum, error) {
	data, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, forum.ErrFailedGetDetailForum
	}
	if data.ID == "" {
		return nil, forum.ErrPageNotFound
	}

	return data, nil
}

func (fu ForumUsecase) Create(createForum *forum.CreateRequest) error {
	category := constant.TOPICS[createForum.Category]
	newForum := entity.Forum{
		ID:       createForum.ID,
		UserId:   createForum.UserId,
		Category: category,
		Link:     createForum.Link,
		Topic:    createForum.Topic,
	}

	err := fu.ForumR.Create(&newForum)
	if err != nil {
		return forum.ErrFailedCreateForum
	}
	return nil
}

func (fu ForumUsecase) Update(id, user_id string, forumId *forum.UpdateRequest) error {
	data, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return forum.ErrFailedGetDetailForum
	}
	if data.ID == "" {
		return forum.ErrPageNotFound
	}

	category := constant.TOPICS[forumId.Category]
	newForum := entity.Forum{
		Category: category,
		Link:     forumId.Link,
		Topic:    forumId.Topic,
	}
	err = fu.ForumR.Update(id, user_id, &newForum)

	if err != nil {
		return forum.ErrFailedUpdateForum
	}
	return nil
}

func (fu ForumUsecase) Delete(id, user_id string) error {
	data, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return forum.ErrFailedGetDetailForum
	}
	if data.ID == "" {
		return forum.ErrPageNotFound
	}

	err = fu.ForumR.Delete(id, user_id)

	if err != nil {
		return forum.ErrFailedDeleteForum
	}
	return nil
}
