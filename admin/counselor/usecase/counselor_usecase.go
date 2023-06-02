package usecase

import (
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CounselorUsecase interface {
	GetAll(offset, limit int) ([]counselor.GetAllResponse, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (counselor.GetByResponse, error)
	Create(inputDetail counselor.CreateRequest, inputProfilePicture *multipart.FileHeader) error
	Update(inputDetail counselor.UpdateRequest, inputProfilePicture *multipart.FileHeader) error
	Delete(id string) error
}

type counselorUsecase struct {
	counselorRepo repository.CounselorRepository	
}

func NewCounselorUsecase(CRepo repository.CounselorRepository) CounselorUsecase {
	return &counselorUsecase{counselorRepo: CRepo}
}

func(u *counselorUsecase) GetAll(offset, limit int) ([]counselor.GetAllResponse, error) {
	
	counselors, err := u.counselorRepo.GetAll(offset, limit)

	if err != nil {
		return nil, counselor.ErrInternalServerError
	}

	return counselors, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.counselorRepo.Count()
	if err != nil {
		return 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) GetById(id string) (counselor.GetByResponse, error) {
	
	counselorRes, err := u.counselorRepo.GetById(id)

	if err != nil {
		return counselorRes, counselor.ErrReviewNotFound
	}

	return counselorRes, nil
}

func(u *counselorUsecase) Create(inputDetail counselor.CreateRequest, inputProfilePicture *multipart.FileHeader) error{
	
	_, err := u.counselorRepo.GetByEmail(inputDetail.Email)

	if err == nil {
		return counselor.ErrCounselorConflict
	}

	path, err := helper.UploadImageToS3(inputProfilePicture)

	if err != nil {
		return counselor.ErrInternalServerError
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()
	
	newCounselor := entity.Counselor{
		ID: uuid,
		Name: inputDetail.Name,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: constant.TOPICS[inputDetail.Topic],
		Description: inputDetail.Description,
		Tarif: inputDetail.Tarif,
		ProfilePicture: path,
	}

	err = u.counselorRepo.Create(newCounselor)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Update(inputDetail counselor.UpdateRequest, inputProfilePicture *multipart.FileHeader) error {

	counselorData, err := u.counselorRepo.GetById(inputDetail.ID)
	
	if err != nil {
		return counselor.ErrReviewNotFound
	}

	_, err = u.counselorRepo.GetByEmail(inputDetail.Email)
	
	if err == nil {	
		if counselorData.Email != inputDetail.Email {
			return counselor.ErrEmailConflict
		}
	}

	counselorUpdate := entity.Counselor{
		Name: inputDetail.Name,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: constant.TOPICS[inputDetail.Topic],
		Description: inputDetail.Description,
		Tarif: inputDetail.Tarif,
	}

	if inputProfilePicture != nil {
		err := helper.DeleteImageFromS3(counselorData.ProfilePicture)

		if err != nil {
			return counselor.ErrInternalServerError
		}
	
		path, err := helper.UploadImageToS3(inputProfilePicture)
		
		if err != nil {
			return counselor.ErrInternalServerError
		}
		
		counselorUpdate.ProfilePicture = path

	}
	
	err = u.counselorRepo.Update(counselorData.ID, counselorUpdate)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Delete(id string) error {
	
	counselorData, err := u.counselorRepo.GetById(id)

	if err != nil {
		return counselor.ErrReviewNotFound
	}
	
	err = u.counselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}
