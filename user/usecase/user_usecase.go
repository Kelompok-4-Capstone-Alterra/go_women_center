package usecase

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/repository"
)

type UserUsecase interface {
	Register(userDTO user.RegisterUserDTO) (error)
	VerifyEmail(email string) error
	
}

type userUsecase struct {
	repo          repository.UserRepo
	UuidGenerator helper.UuidGenerator
	EmailSender   helper.EmailSender
	otpRepo repository.LocalCache
}

func NewUserUsecase(repo repository.UserRepo, idGenerator helper.UuidGenerator, emailSender helper.EmailSender, otpRepo repository.LocalCache) *userUsecase {
	return &userUsecase{
		repo:          repo,
		UuidGenerator: idGenerator,
		EmailSender:   emailSender,
		otpRepo: otpRepo,
	}
}

func (u *userUsecase) Register(userDTO user.RegisterUserDTO) (error) {
	storedOtp, err := u.otpRepo.Read(userDTO.Email)
	if err != nil {
		return err
	}
	
	uuid, err := u.UuidGenerator.GenerateUUID()
	if err != nil {
		return err
	}
	
	defer u.otpRepo.Delete(storedOtp.Email)

	data := domain.User{
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
		Code: otpCode,
	}

	err = u.EmailSender.SendEmail(email, "OTP verification code", otpCode) //TODO: write subject and body template
	if err != nil {
		return err
	}

	u.otpRepo.Update(otp, time.Now().Add(time.Minute).Unix())
	return nil
}
