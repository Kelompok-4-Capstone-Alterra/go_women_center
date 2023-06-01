package usecase

import "github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"

type UserForumUsecaseInterface interface {
	Create(forum *entity.UserForum) (*entity.UserForum, error)
}

type UserForumUsecase struct {
	UserForumR UserForumUsecaseInterface
}

func NewUserForumUsecase(UserForumR UserForumUsecaseInterface) UserForumUsecaseInterface {
	return &UserForumUsecase{
		UserForumR: UserForumR,
	}
}

func (fu UserForumUsecase) Create(forum *entity.UserForum) (*entity.UserForum, error) {
	forum, err := fu.UserForumR.Create(forum)
	if err != nil {
		return nil, err
	}
	return forum, nil
}
