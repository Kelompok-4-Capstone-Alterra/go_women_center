package usecase

import (
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
)

type ForumUsecaseInterface interface {
	GetAll(id_user, topic, popular, created, categories, getMyForum string, offset, limit int) ([]response.ResponseForum, int, error)
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

func (fu ForumUsecase) GetAll(id_user, topic, popular, created, categories, myforum string, offset, limit int) ([]response.ResponseForum, int, error) {
	var forums []response.ResponseForum
	var err error
	var totalData int64

	idCategories, _ := strconv.Atoi(categories)
	categories = constant.TOPICS[idCategories][0]

	if created == "asc" || created == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByCreated(id_user, topic, created, categories, myforum, offset, limit)
	} else if popular == "asc" || popular == "desc" {
		forums, totalData, err = fu.ForumR.GetAllByPopular(id_user, topic, popular, categories, myforum, offset, limit)
	} else {
		forums, totalData, err = fu.ForumR.GetAll(id_user, topic, categories, myforum, offset, limit)
	}

	if err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return forums, totalPages, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*response.ResponseForum, error) {
	forum, err := fu.ForumR.GetById(id, user_id)

	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (fu ForumUsecase) Create(forum *entity.Forum) error {
	topic, _ := strconv.Atoi(forum.Category)
	forum.Category = constant.TOPICS[topic][0]

	err := fu.ForumR.Create(forum)
	if err != nil {
		return err
	}
	return nil
}

func (fu ForumUsecase) Update(id, user_id string, forumId *entity.Forum) error {
	err := fu.ForumR.Update(id, user_id, forumId)

	if err != nil {
		return err
	}
	return nil
}

func (fu ForumUsecase) Delete(id, user_id string) error {
	err := fu.ForumR.Delete(id, user_id)

	if err != nil {
		return err
	}
	return nil
}
