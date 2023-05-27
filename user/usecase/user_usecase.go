package usecase

import (
	"errors"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/repository"
)

type UserUsecase interface {
	Register(userDTO user.RegisterUserDTO) (domain.User, error)
	VerifyEmail(email string) error
}

type userUsecase struct {
	repo          repository.UserRepo
	UuidGenerator helper.UuidGenerator
	EmailSender   helper.EmailSender
}

func NewUserUsecase(repo repository.UserRepo, idGenerator helper.UuidGenerator, emailSender helper.EmailSender) *userUsecase {
	return &userUsecase{
		repo:          repo,
		UuidGenerator: idGenerator,
		EmailSender:   emailSender,
	}
}

func (u *userUsecase) Register(userDTO user.RegisterUserDTO) (domain.User, error) {
	storedOtp, notEmpty := otpCache[userDTO.Email]
	if !notEmpty {
		return domain.User{}, errors.New("otp is not registed")
	}

	timeSinceOtpSent := time.Since(storedOtp.Deadline).Minutes()
	if timeSinceOtpSent > 1 {
		return domain.User{}, errors.New("invalid or past expirations otp")
	}

	defer delete(otpCache, userDTO.Email)

	uuid, err := u.UuidGenerator.GenerateUUID()
	if err != nil {
		return domain.User{}, err
	}

	data := domain.User{
		ID:       uuid,
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Username: userDTO.Username,
		Password: userDTO.Password,
	}

	return u.repo.Create(data)
}

func (u *userUsecase) VerifyEmail(email string) error {
	otpCode, err := helper.GetOtp()
	if err != nil {
		return err
	}

	err = u.EmailSender.SendEmail(email, "OTP verification code", otpCode) //TODO: write subject and body template
	if err != nil {
		return err
	}

	otpCache[email] = domain.NewOTP(otpCode)
	return nil
}

var otpCache = map[string]domain.OTP{}
