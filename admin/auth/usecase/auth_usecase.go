package usecase

import (
	"fmt"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type AuthUsecase interface {
	Login(request auth.LoginAdminRequest) (entity.Admin, error)
}

type authUseCase struct {
	Repo repository.AdminRepo
	Encryptor     helper.Encryptor
}

func NewAuthUsecase(repo repository.AdminRepo, encryptor helper.Encryptor) AuthUsecase {
	return &authUseCase{
		Repo: repo,
		Encryptor: encryptor,
	}
}

func (u *authUseCase) Login(request auth.LoginAdminRequest) (entity.Admin, error) {
	data, err := u.Repo.GetByUsername(request.Username)
	if err != nil {
		return entity.Admin{}, auth.ErrInternalServerError
	}
	fmt.Println(request.Password + " " + data.Password)
	if !u.Encryptor.CheckPasswordHash(request.Password, data.Password) {
		return entity.Admin{}, auth.ErrInvalidCredential
	}

	return data, nil
}
