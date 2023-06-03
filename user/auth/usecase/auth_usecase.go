package usecase

import (
	"log"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	user "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
)

type UserUsecase interface {
	Register(registerRequest user.RegisterUserRequest) error
	VerifyEmail(email string) error
	Login(loginRequest user.LoginUserRequest) (entity.User, error)
}

type userUsecase struct {
	repo          repository.UserRepo
	UuidGenerator helper.UuidGenerator
	EmailSender   helper.EmailSender
	otpRepo       repository.LocalCache
	otpGen        helper.OtpGenerator
	Encryptor     helper.Encryptor
}

func NewUserUsecase(repo repository.UserRepo, idGenerator helper.UuidGenerator, emailSender helper.EmailSender, otpRepo repository.LocalCache, otpgen helper.OtpGenerator, encryptor helper.Encryptor) *userUsecase {
	return &userUsecase{
		repo:          repo,
		UuidGenerator: idGenerator,
		EmailSender:   emailSender,
		otpRepo:       otpRepo,
		otpGen:        otpgen,
		Encryptor:     encryptor,
	}
}

func (u *userUsecase) Register(registerRequest user.RegisterUserRequest) error {
	storedOtp, err := u.otpRepo.Read(registerRequest.Email)
	if err != nil {
		return err
	}

	uuid, err := u.UuidGenerator.GenerateUUID()
	if err != nil {
		return err
	}

	
	encryptedPass, err := u.Encryptor.HashPassword(registerRequest.Password)
	if err != nil {
		return user.ErrFailedEncrpyt
	}
	log.Println(registerRequest.Password)
	log.Println(encryptedPass)

	defer u.otpRepo.Delete(storedOtp.Email)

	data := entity.User{
		ID:       uuid,
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Username: registerRequest.Username,
		Password: encryptedPass,
	}

	_, err = u.repo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) VerifyEmail(email string) error {
	otpCode, err := u.otpGen.GetOtp()
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

func (u *userUsecase) Login(loginRequest user.LoginUserRequest) (entity.User, error) {
	data, err := u.repo.GetByEmail(loginRequest.Email)
	if err != nil {
		return entity.User{}, user.ErrInvalidCredential
	}

	if !u.Encryptor.CheckPasswordHash(loginRequest.Password, data.Password) {
		return entity.User{}, user.ErrInvalidCredential
	}

	return data, nil
}
