package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/repository"
)

type UserUsecase interface {
	Register(userDTO user.RegisterUserDTO) (domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepo
}

func NewUserUsecase(repo repository.UserRepo) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Register(userDTO user.RegisterUserDTO) (domain.User, error) {
	data := domain.User{
		Name: userDTO.Name,
		Email: userDTO.Email,
		Username: userDTO.Username,
		Password: userDTO.Password,
	}
	return u.repo.Create(data)
}