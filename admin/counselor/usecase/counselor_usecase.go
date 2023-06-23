package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	repoSchedule "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type CounselorUsecase interface {
	GetAll(search, sortBy, hasSchedule string, offset, limit int) ([]counselor.GetAllResponse, int, error)
	GetById(id string) (counselor.GetByResponse, error)
	Create(input counselor.CreateRequest) error
	Update(input counselor.UpdateRequest) error
	Delete(id string) error	
}

type counselorUsecase struct {
	CounselorRepo repo.CounselorRepository
	ScheduleRepo repoSchedule.ScheduleRepository
	Image helper.Image
	uuidGenerator helper.UuidGenerator
}

func NewCounselorUsecase(CRepo repo.CounselorRepository, SRepo repoSchedule.ScheduleRepository, Image helper.Image, UuidGenerator helper.UuidGenerator) CounselorUsecase {
	return &counselorUsecase{CRepo, SRepo, Image, UuidGenerator}
}

func(u *counselorUsecase) GetAll(search, sortBy, hasSchedule string, offset, limit int) ([]counselor.GetAllResponse, int, error) {
	
	switch sortBy {
		case "oldest":
			sortBy = "created_at ASC"
		case "newest":
			sortBy = "created_at DESC"
	}
	var counselors, totalData, err = []counselor.GetAllResponse{}, int64(0), error(nil)
	switch hasSchedule {
		case "true":
			counselors, totalData, err = u.CounselorRepo.GetAllHasSchedule(search, sortBy, offset, limit)
		case "false":
			counselors, totalData, err = u.CounselorRepo.GetAllNotHasSchedule(search, sortBy, offset, limit)
		default:
			counselors, totalData, err = u.CounselorRepo.GetAll(search, sortBy, offset, limit)
	}

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

	uuid, err := u.uuidGenerator.GenerateUUID()

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	
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
	
	if counselorData.ProfilePicture != "" {

		err = u.Image.DeleteImageFromS3(counselorData.ProfilePicture)
	
		if err != nil {
			return counselor.ErrInternalServerError
		}
	}
	
	err = u.CounselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}
