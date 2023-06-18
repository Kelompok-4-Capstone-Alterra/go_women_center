package usecase

import (
	"errors"
	"fmt"

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

	var newCategory string
	category, ok := constant.TOPICS[getAllRequest.Category]
	if ok {
		newCategory = category[0]
	}

	// if created == "asc" || created == "desc" {
	// 	forums, totalData, err = fu.ForumR.GetAllByCreated(id_user, topic, created, categories, myforum, offset, limit)
	// } else if popular == "asc" || popular == "desc" {
	// 	forums, totalData, err = fu.ForumR.GetAllByPopular(id_user, topic, popular, categories, myforum, offset, limit)
	// } else {
	// 	forums, totalData, err = fu.ForumR.GetAll(id_user, topic, categories, myforum, offset, limit)
	// }

	if getAllRequest.SortBy != "" {
		forums, totalData, err = fu.ForumR.GetAllSortBy(getAllRequest, newCategory)
		fmt.Println("masuk if")
	} else {
		fmt.Println("masuk else")
		forums, totalData, err = fu.ForumR.GetAll(getAllRequest, newCategory)
	}

	if err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllRequest.Limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*forum.ResponseForum, error) {
	forum, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, err
	} else if forum.ID == "" {
		return nil, errors.New("invallid id")
	}
	return forum, nil
}

func (fu ForumUsecase) Create(createRequest *forum.CreateRequest) error {
	var newCategory string
	category, ok := constant.TOPICS[createRequest.Category]
	if ok {
		newCategory = category[0]
	}

	if newCategory == "" {
		return errors.New("invalllid category forum")
	}

	forum := entity.Forum{
		ID:       createRequest.ID,
		UserId:   createRequest.UserId,
		Category: newCategory,
		Link:     createRequest.Link,
		Topic:    createRequest.Topic,
		Status:   createRequest.Status,
		Member:   createRequest.Member,
	}

	err := fu.ForumR.Create(&forum)
	if err != nil {
		return err
	}
	return nil
}

func (fu ForumUsecase) Update(id, user_id string, updateRequest *forum.UpdateRequest) error {
	forum, err := fu.ForumR.GetById(id, user_id)
	if err != nil {
		return err
	} else if forum.ID == "" {
		return errors.New("invallid id" + id)
	} else if forum.UserId != user_id {
		return errors.New("cannot access")
	}

	var newCategory string
	category, ok := constant.TOPICS[updateRequest.Category]
	if ok {
		newCategory = category[0]
	}

	if newCategory == "" {
		return errors.New("invalllid category forum")
	}

	updateForum := entity.Forum{
		Category: newCategory,
		Link:     updateRequest.Link,
		Topic:    updateRequest.Topic,
	}
	err = fu.ForumR.Update(id, user_id, &updateForum)

	if err != nil {
		return err
	} else if forum.ID == "" {
		return errors.New("invallid id" + id)
	}
	return nil
}

func (fu ForumUsecase) Delete(id, user_id string) error {
	forum, err := fu.ForumR.GetById(id, user_id)
	if err != nil {
		return err
	} else if forum.ID == "" {
		return errors.New("invallid id " + id)
	}

	err = fu.ForumR.Delete(id, user_id)

	if err != nil {
		return err
	}
	return nil
}
