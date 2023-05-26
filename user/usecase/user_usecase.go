package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/repository"
)

type UserUsecase interface {
	Register(userDTO user.RegisterUserDTO) (domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepo
	UuidGenerator helper.UuidGenerator
}

func NewUserUsecase(repo repository.UserRepo, idGenerator helper.UuidGenerator) *userUsecase {
	return &userUsecase{
		repo: repo,
		UuidGenerator: idGenerator,
	}
}

func (u *userUsecase) Register(userDTO user.RegisterUserDTO) (domain.User, error) {
	uuid, err := u.UuidGenerator.GenerateUUID()
	if err != nil {
		return domain.User{}, err
	}

	data := domain.User{
		ID: uuid,
		Name: userDTO.Name,
		Email: userDTO.Email,
		Username: userDTO.Username,
		Password: userDTO.Password,
	}

	return u.repo.Create(data)
}