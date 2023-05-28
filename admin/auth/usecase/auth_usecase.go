package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
)

type AuthUsecase interface {
	Login(reqDTO auth.LoginAdminDTO) (domain.Admin, error)
}

type authUseCase struct {
	Repo repository.AdminRepo
}

func NewAuthUsecase(repo repository.AdminRepo) *authUseCase {
	return &authUseCase{
		Repo: repo,
	}
}

func (a *authUseCase) Login(reqDTO auth.LoginAdminDTO) (domain.Admin, error) {
	data, err := a.Repo.GetByEmail(reqDTO.Email)
	if err != nil {
		return domain.Admin{}, constant.ErrInvalidCredential
	}

	if reqDTO.Password != data.Password {
		return domain.Admin{}, constant.ErrInvalidCredential
	}

	return data, nil
}