package usecase

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	user "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
)

type UserUsecase interface {
	Register(userDTO user.RegisterUserRequest) error
	VerifyEmail(email string) error
	Login(userDTO user.LoginUserRequest) (entity.User, error)
}

type userUsecase struct {
	repo          repository.UserRepo
	UuidGenerator helper.UuidGenerator
	EmailSender   helper.EmailSender
	otpRepo       repository.LocalCache
}

func NewUserUsecase(repo repository.UserRepo, idGenerator helper.UuidGenerator, emailSender helper.EmailSender, otpRepo repository.LocalCache) *userUsecase {
	return &userUsecase{
		repo:          repo,
		UuidGenerator: idGenerator,
		EmailSender:   emailSender,
		otpRepo:       otpRepo,
	}
}

func (u *userUsecase) Register(userDTO user.RegisterUserRequest) error {
	storedOtp, err := u.otpRepo.Read(userDTO.Email)
	if err != nil {
		return err
	}

	uuid, err := u.UuidGenerator.GenerateUUID()
	if err != nil {
		return err
	}

	defer u.otpRepo.Delete(storedOtp.Email)

	data := entity.User{
		ID:       uuid,
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Username: userDTO.Username,
		Password: userDTO.Password,
	}

	_, err = u.repo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) VerifyEmail(email string) error {
	otpCode, err := helper.GetOtp()
	if err != nil {
		return err
	}
	otp := repository.Otp{
		Email: email,
		Code:  otpCode,
	}

	err = u.EmailSender.SendEmail(email, "OTP verification code (valid for 1 minute)", otpCode) //TODO: write subject and body template
	if err != nil {
		return err
	}

	u.otpRepo.Update(otp, time.Now().Add(time.Minute).Unix())
	return nil
}

func (u *userUsecase) Login(userDTO user.LoginUserRequest) (entity.User, error) {
	data, err := u.repo.GetByEmail(userDTO.Email)
	if err != nil {
		return entity.User{}, constant.ErrInvalidCredential
	}

	if userDTO.Password != data.Password {
		return entity.User{}, constant.ErrInvalidCredential
	}

	return data, nil
}
