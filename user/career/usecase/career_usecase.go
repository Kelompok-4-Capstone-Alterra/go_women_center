package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/repository"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CareerUsecase interface {
	GetAll(offset, limit int) ([]career.GetAllResponse, error)
	GetById(id string) (career.GetByResponse, error)
	GetBySearch(search string) ([]career.GetAllResponse, error)
	GetTotalPages(limit int) (int, error)
}

type careerUsecase struct {
	careerRepo repository.CareerRepository
}

func NewCareerUsecase(CRepo repository.CareerRepository) CareerUsecase {
	return &careerUsecase{careerRepo: CRepo}
}

func (u *careerUsecase) GetAll(offset, limit int) ([]career.GetAllResponse, error) {

	careers, err := u.careerRepo.GetAll(offset, limit)

	if err != nil {
		return nil, career.ErrInternalServerError
	}

	return careers, nil
}

func (u *careerUsecase) GetById(id string) (career.GetByResponse, error) {

	careerData, err := u.careerRepo.GetById(id)

	if err != nil {
		return careerData, career.ErrCareerNotFound
	}

	return careerData, nil
}

func (u *careerUsecase) GetBySearch(search string) ([]career.GetAllResponse, error) {

	careers, err := u.careerRepo.GetBySearch(search)

	if err != nil {
		return nil, career.ErrInternalServerError
	}

	return careers, nil
}

func (u *careerUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.careerRepo.Count()
	if err != nil {
		return 0, career.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}
