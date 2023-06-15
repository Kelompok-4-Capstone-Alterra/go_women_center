package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type VoucherUsecase interface {
	GetAll(userId string) ([]entity.Voucher, error)
}

type voucherUsecase struct {
	repo repository.MysqlVoucherRepository
}

func NewtransactionUsecase(
	voucherRepo repository.MysqlVoucherRepository,
) VoucherUsecase {
	return &voucherUsecase{
		repo: voucherRepo,
	}
}

func (u *voucherUsecase) GetAll(userId string) ([]entity.Voucher, error)  {
	data, err := u.repo.GetAll(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}