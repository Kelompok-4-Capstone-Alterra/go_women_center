package usecase

import (
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)


type counselorUsecase struct {
	counselorRepo domain.CounselorRepository
}

func NewCounselorUsecase(counselorRepo domain.CounselorRepository) domain.CounselorUsecase {
	return &counselorUsecase{counselorRepo: counselorRepo}
}

func(u *counselorUsecase) GetAll(offset, limit int) ([]domain.Counselor, error) {
	
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

func(u *counselorUsecase) GetById(id string) (domain.Counselor, error) {
	
	counselorData, err := u.counselorRepo.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return counselorData, counselor.ErrNotFound
		}
		return counselorData, counselor.ErrInternalServerError
	}

	return counselorData, nil
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
	
	newCounselor := domain.Counselor{
		ID: uuid,
		FullName: inputDetail.FullName,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: inputDetail.Topic,
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
		if err.Error() == "record not found" {
			return counselor.ErrNotFound
		}
		return counselor.ErrInternalServerError
	}

	_, err = u.counselorRepo.GetByEmail(inputDetail.Email)
	
	if err == nil {	
		if counselorData.Email != inputDetail.Email {
			return counselor.ErrEmailConflict
		}
	}

	counselorUpdate := domain.Counselor{
		FullName: inputDetail.FullName,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: inputDetail.Topic,
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
		if err.Error() == "record not found" {
			return counselor.ErrNotFound
		}
		return counselor.ErrInternalServerError
	}
	
	err = u.counselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}
