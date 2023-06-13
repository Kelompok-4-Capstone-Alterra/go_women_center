package usecase

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ForumUsecaseInterface interface {
	GetAll(getAllParam forum.GetAllRequest) ([]response.ResponseForum, int, error)
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

func (fu ForumUsecase) GetAll(getAllParam forum.GetAllRequest) ([]response.ResponseForum, int, error) {
	var forums []response.ResponseForum
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
		return nil, 0, errors.New("failed to get all forum data")
	}

	for i := 0; i < len(forums); i++ {
		forums[i].UserForums = nil
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllParam.Limit)

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
