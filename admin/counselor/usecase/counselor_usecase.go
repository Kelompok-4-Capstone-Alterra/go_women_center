package usecase

import (
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
	Create(input counselor.CreateRequest) error
	Update(input counselor.UpdateRequest) error
	Delete(id string) error
	GetTotalPagesSearch(search string, limit int) (int, error)
	Search(search string, offset, limit int) ([]counselor.GetAllResponse, error)
}

type counselorUsecase struct {
	CounselorRepo repository.CounselorRepository
	Image helper.Image
}

func NewCounselorUsecase(CRepo repository.CounselorRepository, Image helper.Image) CounselorUsecase {
	return &counselorUsecase{CounselorRepo: CRepo, Image: Image}
}

func(u *counselorUsecase) GetAll(offset, limit int) ([]counselor.GetAllResponse, error) {
	
	counselors, err := u.CounselorRepo.GetAll(offset, limit)

	if err != nil {
		return nil, counselor.ErrInternalServerError
	}

	return counselors, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.CounselorRepo.Count()
	if err != nil {
		return 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) GetById(id string) (counselor.GetByResponse, error) {
	
	counselorRes, err := u.CounselorRepo.GetById(id)

	if err != nil {
		return counselorRes, counselor.ErrReviewNotFound
	}

	return counselorRes, nil
}

func(u *counselorUsecase) Create(input counselor.CreateRequest) error{
	
	_, err := u.CounselorRepo.GetByEmail(input.Email)

	if err == nil {
		return counselor.ErrCounselorConflict
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
		Topic: constant.TOPICS[input.Topic],
		Description: input.Description,
		Tarif: input.Tarif,
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
		return counselor.ErrReviewNotFound
	}

	_, err = u.CounselorRepo.GetByEmail(input.Email)
	
	if err == nil {	
		if counselorData.Email != input.Email {
			return counselor.ErrEmailConflict
		}
	}

	counselorUpdate := entity.Counselor{
		Name: input.Name,
		Email: input.Email,
		Username: input.Username,
		Topic: constant.TOPICS[input.Topic],
		Description: input.Description,
		Tarif: input.Tarif,
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
		return counselor.ErrReviewNotFound
	}
	
	err = u.CounselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}

func(u *counselorUsecase) GetTotalPagesSearch(search string, limit int) (int, error) {
	
	totalData, err := u.CounselorRepo.CountBySearch(search)
	if err != nil {
		return 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) Search(search string, offset, limit int) ([]counselor.GetAllResponse, error) {
	
	counselors, err := u.CounselorRepo.Search(search, offset, limit)

	if err != nil {
		return nil, counselor.ErrInternalServerError
	}

	return counselors, nil
}


