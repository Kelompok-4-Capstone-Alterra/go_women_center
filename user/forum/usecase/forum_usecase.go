package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
	userForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum"
	repositoryUserForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
	"github.com/google/uuid"
)

type ForumUsecaseInterface interface {
	GetAll(getAllRequest forum.GetAllRequest) ([]forum.ResponseForum, int, error)
	GetById(id, user_id string) (*forum.ResponseForum, error)
	Create(createRequest *forum.CreateRequest) error
	Update(id, user_id string, updateRequest *forum.UpdateRequest) error
	Delete(id, user_id string) error
}

type ForumUsecase struct {
	ForumR     repository.ForumRepository
	UserForumR repositoryUserForum.UserForumRepository
}

func NewForumUsecase(ForumR repository.ForumRepository, UserForumR repositoryUserForum.UserForumRepository) ForumUsecaseInterface {
	return &ForumUsecase{
		ForumR:     ForumR,
		UserForumR: UserForumR,
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
		return nil, 0, forum.ErrFailedGetForum
	}

	if getAllRequest.UserId != "" {
		for i := 0; i < len(forums); i++ {
			for j := 0; j < len(forums[i].UserForums); j++ {
				if forums[i].UserForums[j].UserId == getAllRequest.UserId {
					forums[i].Status = true
					break
				}
			}
		}
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllRequest.Limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*forum.ResponseForum, error) {
	forumId, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, forum.ErrFailedGetDetailForum
	} else if forumId.ID == "" {
		return nil, forum.ErrInvalidId
	}

	if user_id != "" {
		for i := 0; i < len(forumId.UserForums); i++ {
			if forumId.UserForums[i].UserId == user_id {
				forumId.Status = true
				break
			}
		}
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
		return forum.ErrFailedCreateForum
	}

	uuidWithHyphen := uuid.New()
	createUserForum := entity.UserForum{
		ID:      uuidWithHyphen.String(),
		UserId:  createForum.UserId,
		ForumId: createForum.ID,
	}
	err = fu.UserForumR.Create(&createUserForum)
	if err != nil {
		return userForum.ErrFailedCreateUserForum
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
		return forum.ErrFailedUpdateForum
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
		return forum.ErrFailedDeleteForum
	}
	return nil
}
