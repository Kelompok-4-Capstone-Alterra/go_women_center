package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
)

type ForumUsecaseInterface interface {
	GetAll(id_user, topic, popular, created, categories, getMyForum string) ([]response.ResponseForum, error)
	GetById(id, user_id string) (*response.ResponseForum, error)
	Create(forum *entity.Forum) error
	Update(id string, forumId *entity.Forum) error
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

func (fu ForumUsecase) GetAll(id_user, topic, popular, created, categories, myforum string) ([]response.ResponseForum, error) {
	var forums []response.ResponseForum
	var err error

	if created == "asc" || created == "desc" {
		forums, err = fu.ForumR.GetAllByCreated(id_user, topic, created, categories, myforum)
	} else if popular == "desc" {
		forums, err = fu.ForumR.GetAllByPopular(id_user, topic, popular, categories, myforum)
	} else {
		forums, err = fu.ForumR.GetAll(id_user, topic, categories, myforum)
	}

	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetById(id, user_id string) (*response.ResponseForum, error) {
	forum, err := fu.ForumR.GetById(id, user_id)

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
