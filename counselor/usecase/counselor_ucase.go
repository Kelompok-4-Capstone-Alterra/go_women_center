package usecase

import (
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

func(u *counselorUsecase) Create(input counselor.CreateRequest) error {
	
	uuid, _ := helper.NewGoogleUUID().GenerateUUID()
	
	newCounselor := domain.Counselor{
		ID: uuid,
		FullName: input.FullName,
		Email: input.Email,
		Username: input.Username,
		Topic: input.Topic,
		Description: input.Description,
		Tarif: input.Tarif,
		ProfilePicture: input.ProfilePicture,
	}

	err := u.counselorRepo.Create(newCounselor)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}