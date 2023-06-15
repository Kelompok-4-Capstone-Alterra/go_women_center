package usecase

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/repository"
)

type ProfileUsecase interface {
	GetById(id string) (profile.GetByIdResponse, error)
	Update(input profile.UpdateRequest) error
	UpdatePassword(input profile.UpdatePasswordRequest) error
}

type profileUsecase struct {
	profileRepo repo.UserRepository
	image helper.Image
	encryptor helper.Encryptor
}

func NewProfileUsecase(profileRepo repo.UserRepository, image helper.Image, encryptor helper.Encryptor) ProfileUsecase {
	return &profileUsecase{profileRepo, image, encryptor}
}

func(u *profileUsecase) GetById(id string) (profile.GetByIdResponse, error) {
	userData, err := u.profileRepo.GetById(id)

	if err != nil {
		return profile.GetByIdResponse{}, profile.ErrInternalServerError
	}

	profile := profile.GetByIdResponse{
		ID: userData.ID,
		ProfilePicture: userData.ProfilePicture,
		Username: userData.Username,
		Name: userData.Name,
		Email: userData.Email,
		PhoneNumber: userData.PhoneNumber,
	}

	return profile, nil
}

func(u *profileUsecase) Update(input profile.UpdateRequest) error {
	

	user, err := u.profileRepo.GetById(input.ID)

	if err != nil {
		log.Println(err.Error())
		return profile.ErrInternalServerError
	}
	
	userProfile := entity.User{
		ID: input.ID,
		Name: input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	if input.Username != "" && (user.Username != input.Username) {
		err := u.profileRepo.GetByUsername(input.Username)		
	
		if err == nil {
			return profile.ErrUsernameDuplicate
		}
		
		userProfile.Username = input.Username
	}

	if input.ProfilePicture != nil {

		
		if !u.image.IsImageValid(input.ProfilePicture) {
			return profile.ErrProfilePictureFormat
		}

		path, err := u.image.UploadImageToS3(input.ProfilePicture)

		if err != nil {
			return profile.ErrInternalServerError
		}

		err = u.image.DeleteImageFromS3(user.ProfilePicture)

		if err != nil {
			return profile.ErrInternalServerError
		}

		userProfile.ProfilePicture = path

	}

	err = u.profileRepo.Update(userProfile)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func(u *profileUsecase) UpdatePassword(input profile.UpdatePasswordRequest) error {
	
	user, err := u.profileRepo.GetById(input.ID)

	if err != nil {
		return profile.ErrInternalServerError
	}

	if !u.encryptor.CheckPasswordHash(input.CurrentPassword, user.Password) {
		return profile.ErrPasswordNotMatch
	}

	user.Password, _ = u.encryptor.HashPassword(input.NewPassword)

	err = u.profileRepo.Update(user)

	if err != nil {
		return profile.ErrInternalServerError
	}

	return nil
}