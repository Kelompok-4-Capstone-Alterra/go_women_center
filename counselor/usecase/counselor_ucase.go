package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)


type counselorUsecase struct {
	counselorRepo domain.CounselorRepository
}

func NewCounselorUsecase(counselorRepo domain.CounselorRepository) domain.CounselorUsecase {
	return &counselorUsecase{counselorRepo: counselorRepo}
}

func(u *counselorUsecase) GetAll(page, limit int) ([]domain.Counselor, error) {
	
	offset, limit := helper.GetOffsetAndLimit(page, limit)
	
	counselors, err := u.counselorRepo.GetAll(offset, limit)

	if err != nil {
		return nil, domain.ErrInternalServerError
	}

	return counselors, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.counselorRepo.Count()
	if err != nil {
		return 0, domain.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) Create(counselor domain.Counselor) error {

	

	err := u.counselorRepo.Create(counselor)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}