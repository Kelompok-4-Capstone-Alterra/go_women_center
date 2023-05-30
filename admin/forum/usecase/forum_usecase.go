package usecase

import "github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"

type ForumUsecase struct {
	ForumR entity.ForumRepository
}

func NewForumUsecase(ForumR entity.ForumRepository) entity.ForumUsecase {
	return &ForumUsecase{
		ForumR: ForumR,
	}
}

func (fu ForumUsecase) GetAll() ([]entity.Forum, error) {
	forums, err := fu.ForumR.GetAll()

	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fu ForumUsecase) GetById(id string) (*entity.Forum, error) {
	return &entity.Forum{}, nil
}

func (fu ForumUsecase) Create(forum entity.Forum) (entity.Forum, error) {
	return entity.Forum{}, nil
}

func (fu ForumUsecase) Update(id string, forumId entity.Forum) (entity.Forum, error) {
	return entity.Forum{}, nil
}

func (fu ForumUsecase) Delete(id string) error {
	return nil
}
