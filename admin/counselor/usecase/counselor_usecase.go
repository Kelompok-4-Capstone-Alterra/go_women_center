package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CounselorUsecase interface {
	GetAll(search string, offset, limit int) ([]counselor.GetAllResponse, int, error)
	GetById(id string) (counselor.GetByResponse, error)
	Create(input counselor.CreateRequest) error
	Update(input counselor.UpdateRequest) error
	Delete(id string) error	
}

type counselorUsecase struct {
	CounselorRepo repository.CounselorRepository
	Image helper.Image
}

func NewCounselorUsecase(CRepo repository.CounselorRepository, Image helper.Image) CounselorUsecase {
	return &counselorUsecase{CounselorRepo: CRepo, Image: Image}
}

func(u *counselorUsecase) GetAll(search string, offset, limit int) ([]counselor.GetAllResponse, int, error) {
	
	counselors, totalData, err := u.CounselorRepo.GetAll(search, offset, limit)

	if err != nil {
		return nil, 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(int(totalData),limit)

	return counselors, totalPages , nil
}

func(u *counselorUsecase) GetById(id string) (counselor.GetByResponse, error) {
	
	counselorRes, err := u.CounselorRepo.GetById(id)

	if err != nil {
		return counselorRes, counselor.ErrCounselorNotFound
	}

	return counselorRes, nil
}

func(u *counselorUsecase) Create(input counselor.CreateRequest) error{
	
	_, err := u.CounselorRepo.GetByEmail(input.Email)

	if err == nil {
		return counselor.ErrEmailConflict
	}

	_, err = u.CounselorRepo.GetByUsername(input.Username)

	if err == nil {
		return counselor.ErrUsernameConflict
	}


	if !u.Image.IsImageValid(input.ProfilePicture) {
		return counselor.ErrProfilePictureFormat
	}

	path, err := u.Image.UploadImageToS3(input.ProfilePicture)

	if err != nil {
		return counselor.ErrInternalServerError
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()
	
	newCounselor := entity.Counselor{
		ID: uuid,
		Name: input.Name,
		Email: input.Email,
		Username: input.Username,
		Topic: constant.TOPICS[input.Topic][0],
		Description: input.Description,
		Price: input.Price,
		ProfilePicture: path,
	}

	err = u.CounselorRepo.Create(newCounselor)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Update(input counselor.UpdateRequest) error {

	counselorData, err := u.CounselorRepo.GetById(input.ID)
	
	if err != nil {
		return counselor.ErrCounselorNotFound
	}

	_, err = u.CounselorRepo.GetByEmail(input.Email)
	
	if err == nil {	
		if counselorData.Email != input.Email {
			return counselor.ErrEmailConflict
		}
	}

	_, err = u.CounselorRepo.GetByUsername(input.Username)

	if err == nil {
		if counselorData.Username != input.Username {
			return counselor.ErrUsernameConflict
		}
	}
	
	counselorUpdate := entity.Counselor{
		Name: input.Name,
		Email: input.Email,
		Username: input.Username,
		Description: input.Description,
		Price: input.Price,
	}
	
	if topic, ok := constant.TOPICS[input.Topic]; ok {
		counselorUpdate.Topic = topic[0]
	}

	if input.ProfilePicture != nil {
		err := u.Image.DeleteImageFromS3(counselorData.ProfilePicture)

		if err != nil {
			return counselor.ErrInternalServerError
		}
	
		path, err := u.Image.UploadImageToS3(input.ProfilePicture)
		
		if err != nil {
			return counselor.ErrInternalServerError
		}
		
		counselorUpdate.ProfilePicture = path

	}
	
	err = u.CounselorRepo.Update(counselorData.ID, counselorUpdate)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Delete(id string) error {
	
	counselorData, err := u.CounselorRepo.GetById(id)

	if err != nil {
		return counselor.ErrCounselorNotFound
	}
	
	err = u.CounselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}
