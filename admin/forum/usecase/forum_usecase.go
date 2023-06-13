package usecase

import (
	"errors"
	"strconv"

	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ForumUsecaseInterface interface {
	GetAll(topic, popular, created, categories, getMyForum string, offset, limit int) ([]response.ResponseForum, int, error)
	GetById(id string) (*response.ResponseForum, error)
	Delete(id string) error
}

type ForumUsecase struct {
	ForumR repository.ForumRepository
}

func NewForumUsecase(ForumR repository.ForumRepository) ForumUsecaseInterface {
	return &ForumUsecase{
		ForumR: ForumR,
	}
}

func (fu ForumUsecase) GetAll(topic, popular, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int, error) {
	var forums []response.ResponseForum
	var err error
	var totalData int64

	idCategories, _ := strconv.Atoi(categories)
	categories = constant.TOPICS[idCategories]

	if created == "asc" || created == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByCreated(topic, created, categories, myforum, offset, limit)
	} else if popular == "asc" || popular == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByPopular(topic, popular, categories, myforum, offset, limit)
	} else {
		forums, totalData, err = fu.ForumR.GetAll(topic, categories, myforum, offset, limit)
	}

	if err != nil {
		return nil, 0, errors.New("failed to get all forum data")
	}

	for i := 0; i < len(forums); i++ {
		forums[i].UserForums = nil
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id string) (*response.ResponseForum, error) {
	forum, err := fu.ForumR.GetById(id)

	if err != nil {
		return nil, errors.New("failed to get forum data details")
	} else if forum.ID == "" {
		return nil, errors.New("invalid forum id " + id)
	}

	forum.UserForums = nil
	return forum, nil
}

func (fu ForumUsecase) Delete(id string) error {
	forum, err := fu.ForumR.GetById(id)

	if err != nil {
		return errors.New("failed to get forum data details")
	} else if forum.ID == "" {
		return errors.New("invalid forum id " + id)
	}

	err2 := fu.ForumR.Delete(id)
	if err2 != nil {
		return errors.New("failed to delete forum data")
	}
	return nil
}
