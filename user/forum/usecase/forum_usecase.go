package usecase

import (
	"errors"
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	request "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
)

type ForumUsecaseInterface interface {
	GetAll(getAllParam request.QueryParamRequest) ([]response.ResponseForum, int, error)
	GetById(id, user_id string) (*response.ResponseForum, error)
	Create(forum *entity.Forum) error
	Update(id, user_id string, forumId *entity.Forum) error
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

func (fu ForumUsecase) GetAll(getAllParam request.QueryParamRequest) ([]response.ResponseForum, int, error) {
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

	totalPages := helper.GetTotalPages(int(totalData), getAllParam.Limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*response.ResponseForum, error) {
	forum, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, errors.New("failed to get forum data details")
	}
	if forum.ID == "" {
		return nil, errors.New("invalid forum id " + id)
	}

	return forum, nil
}

func (fu ForumUsecase) Create(forum *entity.Forum) error {
	topic, _ := strconv.Atoi(forum.Category)
	forum.Category = constant.TOPICS[topic]

	err := fu.ForumR.Create(forum)
	if err != nil {
		return errors.New("failed created forum data")
	}
	return nil
}

func (fu ForumUsecase) Update(id, user_id string, forumId *entity.Forum) error {
	forum, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return errors.New("failed to get forum data details")
	}
	if forum.ID == "" {
		return errors.New("page not found")
	}

	err2 := fu.ForumR.Update(id, user_id, forumId)

	if err2 != nil {
		return errors.New("failed to updated forum data")
	}
	return nil
}

func (fu ForumUsecase) Delete(id, user_id string) error {
	forum, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return errors.New("failed to get forum data details")
	}
	if forum.ID == "" {
		return errors.New("page not found")
	}

	err2 := fu.ForumR.Delete(id, user_id)

	if err2 != nil {
		return errors.New("failed to delete forum data")
	}
	return nil
}
