package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type AuthUsecase interface {
	Login(request auth.LoginAdminRequest) (entity.Admin, error)
}

type authUseCase struct {
	Repo repository.AdminRepo
}

func NewAuthUsecase(repo repository.AdminRepo) *authUseCase {
	return &authUseCase{
		Repo: repo,
	}
}

func (a *authUseCase) Login(request auth.LoginAdminRequest) (entity.Admin, error) {
	data, err := a.Repo.GetByEmail(request.Email)
	if err != nil {
		return entity.Admin{}, auth.ErrInvalidCredential
	}

	if request.Password != data.Password {
		return entity.Admin{}, auth.ErrInvalidCredential
	}

	return data, nil
}
