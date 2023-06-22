package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/repository"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CareerUsecase interface {
	GetAll(search, sortBy string, offset, limit int) ([]career.GetAllResponse, int, error)
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

func (u *careerUsecase) GetAll(search, sortBy string, offset, limit int) ([]career.GetAllResponse, int, error) {
	switch sortBy {
	case "newest":
		sortBy = "created_at DESC"
	case "highest_salary":
		sortBy = "min_salary DESC"
	}

	careers, totalData, err := u.careerRepo.GetAll(search, sortBy, offset, limit)

	if err != nil {
		return nil, 0, career.ErrInternalServerError
	}

	return careers, helper.GetTotalPages(int(totalData), limit), nil
}

func (u *careerUsecase) GetById(id string) (career.GetByResponse, error) {

	careerData, err := u.careerRepo.GetById(id)

	careerDataResponse := career.GetByResponse{
		ID:            careerData.ID,
		Image:         careerData.Image,
		JobPosition:   careerData.JobPosition,
		CompanyName:   careerData.CompanyName,
		Location:      careerData.Location,
		Salary:        careerData.Salary,
		MinExperience: careerData.MinExperience,
		LastEducation: careerData.LastEducation,
		CompanyEmail:  careerData.CompanyEmail,
		Description:   careerData.Description,
		CreatedAt:     careerData.CreatedAt.Format("2006-01-02"),
	}

	if err != nil {
		return careerDataResponse, career.ErrCareerNotFound
	}

	return careerDataResponse, nil
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
