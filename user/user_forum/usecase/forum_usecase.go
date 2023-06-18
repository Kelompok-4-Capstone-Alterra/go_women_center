package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	userForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
)

type UserForumUsecaseInterface interface {
	Create(createUserForum *userForum.CreateRequest) error
}

type UserForumUsecase struct {
	UserForumR repository.UserForumRepository
}

func NewUserForumUsecase(UserForumR repository.UserForumRepository) UserForumUsecaseInterface {
	return &UserForumUsecase{
		UserForumR: UserForumR,
	}
}

func (ufu UserForumUsecase) Create(createUserForum *userForum.CreateRequest) error {
	userForum := entity.UserForum{
		ID:      createUserForum.ID,
		UserId:  createUserForum.UserId,
		ForumId: createUserForum.ForumId,
	}

	err := ufu.UserForumR.Create(&userForum)
	if err != nil {
		return err
	}
	return nil
}
