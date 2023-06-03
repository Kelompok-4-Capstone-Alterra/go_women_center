package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type AuthUsecase interface {
	Login(reqDTO auth.LoginAdminRequest) (entity.Admin, error)
}

type authUseCase struct {
	Repo repository.AdminRepo
}

func NewAuthUsecase(repo repository.AdminRepo) *authUseCase {
	return &authUseCase{
		Repo: repo,
	}
}

func (a *authUseCase) Login(reqDTO auth.LoginAdminRequest) (entity.Admin, error) {
	data, err := a.Repo.GetByEmail(reqDTO.Email)
	if err != nil {
		return entity.Admin{}, constant.ErrInvalidCredential
	}

	if reqDTO.Password != data.Password {
		return entity.Admin{}, constant.ErrInvalidCredential
	}

	return data, nil
}
