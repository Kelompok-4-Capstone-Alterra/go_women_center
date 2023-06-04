package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
)

type ForumUsecaseInterface interface {
	GetAll(topic string) ([]response.ResponseForum, error)
	GetAllSortBy(newest string) ([]response.ResponseForum, error)
	GetByCategory(id_category string) ([]response.ResponseForum, error)
	GetByMyForum(id_user string) ([]response.ResponseForum, error)
	GetById(id string) (*response.ResponseForumDetail, error)
	Create(forum *entity.Forum) error
	Update(id string, forumId *entity.Forum) error
	Delete(id string) error
}

type ForumUsecase struct {
	ForumR ForumUsecaseInterface
}

func NewForumUsecase(ForumR ForumUsecaseInterface) ForumUsecaseInterface {
	return &ForumUsecase{
		ForumR: ForumR,
	}
}

func (fu ForumUsecase) GetAll(topic string) ([]response.ResponseForum, error) {
	forums, err := fu.ForumR.GetAll(topic)
	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetAllSortBy(newest string) ([]response.ResponseForum, error) {
	forums, err := fu.ForumR.GetAllSortBy(newest)
	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetByCategory(id_category string) ([]response.ResponseForum, error) {
	forums, err := fu.ForumR.GetByCategory(id_category)
	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetByMyForum(id_user string) ([]response.ResponseForum, error) {
	forums, err := fu.ForumR.GetByMyForum(id_user)
	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetById(id string) (*response.ResponseForumDetail, error) {
	forum, err := fu.ForumR.GetById(id)

	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (fu ForumUsecase) Create(forum *entity.Forum) error {
	err := fu.ForumR.Create(forum)
	if err != nil {
		return err
	}
	return nil
}

func (fu ForumUsecase) Update(id string, forumId *entity.Forum) error {
	err := fu.ForumR.Update(id, forumId)

	if err != nil {
		return err
	}
	return nil
}

func (fu ForumUsecase) Delete(id string) error {
	err := fu.ForumR.Delete(id)

	if err != nil {
		return err
	}
	return nil
}
