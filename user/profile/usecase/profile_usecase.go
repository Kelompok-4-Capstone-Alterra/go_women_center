package usecase

import (
	"log"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/repository"
	"golang.org/x/sync/errgroup"
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

	if userData.BirthDate != nil {
		profile.BirthDate = userData.BirthDate.Format(time.DateOnly)
	}

	return profile, nil
}

func(u *profileUsecase) Update(input profile.UpdateRequest) error {
	
	var g errgroup.Group
	
	var user entity.User

	g.Go(func() error {
		var err error
		user, err = u.profileRepo.GetById(input.ID)
	
		if err != nil {
			return profile.ErrInternalServerError
		}
		return nil
	})

	var birthDate time.Time
	
	g.Go(func() error {
		var err error
		birthDate, err = time.Parse(time.DateOnly, input.BirthDate)
	
		if err != nil {
			return profile.ErrBirthDateFormat
		}
		return nil
	})

	if user.Email != input.Email {
		g.Go(func() error {
			
			err := u.profileRepo.GetByEmail(input.Email)
		
			if err != nil {
				return profile.ErrEmailDuplicate
			}
			return nil
		})
	
		g.Go(func() error {
			err := u.profileRepo.GetByUsername(input.Username)		
		
			if err != nil {
				return profile.ErrUsernameDuplicate
			}
			return nil
		})
	}

	if err := g.Wait(); err !=nil{
		return err
	}

	userProfile := entity.User{
		ID: input.ID,
		Username: input.Username,
		Name: input.Name,
		Email: input.Email,
		PhoneNumber: input.PhoneNumber,
		BirthDate: &birthDate,
	}

	if input.ProfilePicture != nil {

		if !u.image.IsImageValid(input.ProfilePicture) {
			return profile.ErrProfilePictureFormat
		}

		path, err := u.image.UploadImageToS3(input.ProfilePicture)

		if err != nil {
			return profile.ErrInternalServerError
		}
		
		userProfile.ProfilePicture = path

		err = u.image.DeleteImageFromS3(user.ProfilePicture)

		if err != nil {
			return profile.ErrInternalServerError
		}

	}

	err := u.profileRepo.Update(userProfile)

	if err != nil {
		log.Println(err)
		return profile.ErrInternalServerError
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