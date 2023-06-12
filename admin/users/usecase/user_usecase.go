package usecase

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type UserUsecase interface {
	GetById(id string) (users.GetByIdResponse, error)
	GetAll(search string, offset, limit int) ([]users.GetAllResponse, int, error)
	Delete(id string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) GetById(id string) (users.GetByIdResponse, error) {
	userRes, err := u.userRepository.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return userRes, users.ErrUserNotFound
		}
		log.Println(err.Error())
		return userRes, users.ErrInternalServerError
	}

	return userRes, nil
}

func (u *userUsecase) GetAll(search string, offset, limit int) ([]users.GetAllResponse, int,error) {

	usersRes, totalData, err := u.userRepository.GetAll(search, offset, limit)
	
	if err != nil {
		log.Println(err.Error())
		return []users.GetAllResponse{}, 0, users.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return usersRes, totalPages, nil
}

func (u *userUsecase) Delete(id string) error {

	_,err := u.userRepository.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return users.ErrUserNotFound
		}
		log.Println(err.Error())
		return users.ErrInternalServerError
	}
	

	err = u.userRepository.Delete(id)
	
	if err != nil {
		return err
	}

	return nil
}