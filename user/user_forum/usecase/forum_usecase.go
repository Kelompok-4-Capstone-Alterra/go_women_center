package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
)

type UserForumUsecaseInterface interface {
	Create(forum *entity.UserForum) error
}

type UserForumUsecase struct {
	UserForumR repository.UserForumRepository
}

func NewUserForumUsecase(UserForumR repository.UserForumRepository) UserForumUsecaseInterface {
	return &UserForumUsecase{
		UserForumR: UserForumR,
	}
}

func (fu UserForumUsecase) Create(forum *entity.UserForum) error {
	err := fu.UserForumR.Create(forum)
	if err != nil {
		return err
	}
	return nil
}
