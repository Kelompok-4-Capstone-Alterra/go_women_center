package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/repository"
)

type ProfileUsecase interface {
	GetById(id string) (profile.GetByIdResponse, error)
}

type profileUsecase struct {
	profileRepo repo.UserRepository
}

func NewProfileUsecase(profileRepo repo.UserRepository) ProfileUsecase {
	return &profileUsecase{profileRepo}
}

func(u *profileUsecase) GetById(id string) (profile.GetByIdResponse, error) {
	profile, err := u.profileRepo.GetById(id)

	if err != nil {
		return profile, err
	}

	return profile, nil
}