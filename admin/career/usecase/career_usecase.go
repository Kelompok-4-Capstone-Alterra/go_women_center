package usecase

import (
	"log"
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/repository"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CareerUsecase interface {
	GetAll(search string, offset, limit int) ([]career.GetAllResponse, int, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (career.GetByResponse, error)
	Create(inputDetail career.CreateRequest, inputImage *multipart.FileHeader) error
	Update(inputDetail career.UpdateRequest, inputImage *multipart.FileHeader) error
	Delete(id string) error
}

type careerUsecase struct {
	careerRepo repository.CareerRepository
	image      helper.Image
}

func NewCareerUsecase(CRepo repository.CareerRepository, Image helper.Image) CareerUsecase {
	return &careerUsecase{careerRepo: CRepo, image: Image}
}

func (u *careerUsecase) GetAll(search string, offset, limit int) ([]career.GetAllResponse, int, error) {

	careers, totalData, err := u.careerRepo.GetAll(search, offset, limit)

	if err != nil {
		return nil, 0, career.ErrInternalServerError
	}

	return careers, helper.GetTotalPages(int(totalData), limit), nil
}

func (u *careerUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.careerRepo.Count()
	if err != nil {
		return 0, career.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func (u *careerUsecase) GetById(id string) (career.GetByResponse, error) {

	careerData, err := u.careerRepo.GetById(id)

	if err != nil {
		return careerData, career.ErrCareerNotFound
	}

	return careerData, nil
}

func (u *careerUsecase) Create(inputDetail career.CreateRequest, inputImage *multipart.FileHeader) error {
	path, err := u.image.UploadImageToS3(inputImage)

	if err != nil {
		return career.ErrInternalServerError
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()

	newCareer := entity.Career{
		ID:            uuid,
		JobPosition:   inputDetail.JobPosition,
		CompanyName:   inputDetail.CompanyName,
		Location:      inputDetail.Location,
		Salary:        inputDetail.Salary,
		MinExperience: inputDetail.MinExperience,
		LastEducation: inputDetail.LastEducation,
		Description:   inputDetail.Description,
		CompanyEmail:  inputDetail.CompanyEmail,
		Image:         path,
	}

	err = u.careerRepo.Create(newCareer)

	if err != nil {
		return career.ErrInternalServerError
	}

	return nil
}

func (u *careerUsecase) Update(inputDetail career.UpdateRequest, inputImage *multipart.FileHeader) error {

	careerData, err := u.careerRepo.GetById(inputDetail.ID)

	if err != nil {
		return career.ErrCareerNotFound
	}

	careerUpdate := entity.Career{
		JobPosition:   inputDetail.JobPosition,
		CompanyName:   inputDetail.CompanyName,
		Location:      inputDetail.Location,
		Salary:        inputDetail.Salary,
		MinExperience: inputDetail.MinExperience,
		LastEducation: inputDetail.LastEducation,
		Description:   inputDetail.Description,
		CompanyEmail:  inputDetail.CompanyEmail,
	}

	if inputImage != nil {
		err := u.image.DeleteImageFromS3(careerData.Image)

		if err != nil {
			return career.ErrInternalServerError
		}

		path, err := u.image.UploadImageToS3(inputImage)

		if err != nil {
			return career.ErrInternalServerError
		}

		careerUpdate.Image = path

	}

	err = u.careerRepo.Update(careerData.ID, careerUpdate)

	if err != nil {
		return career.ErrInternalServerError
	}

	return nil
}

func (u *careerUsecase) Delete(id string) error {

	careerData, err := u.careerRepo.GetById(id)

	if err != nil {
		return career.ErrCareerNotFound
	}

	err = u.careerRepo.Delete(careerData.ID)

	if careerData.Image != "" {
		err = u.image.DeleteImageFromS3(careerData.Image)
		if err != nil {
			log.Println(err.Error())
			return career.ErrCareerNotFound
		}
	}

	if err != nil {
		return career.ErrInternalServerError
	}

	return nil
}
