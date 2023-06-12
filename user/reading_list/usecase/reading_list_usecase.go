package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/repository"
)

type ReadingListUsecaseInterface interface {
	GetAll(id_user, name string, offset, limit int) ([]response.ReadingList, int, error)
	GetById(id, user_id string) (*response.ReadingList, error)
	Create(readingList *entity.ReadingList) error
	Update(id, user_id string, readingList *entity.ReadingList) error
	Delete(id, user_id string) error
}

type ReadingListUsecase struct {
	ReadingListR repository.ReadingListRepository
}

func NewReadingListUsecase(ReadingListR repository.ReadingListRepository) ReadingListUsecaseInterface {
	return &ReadingListUsecase{
		ReadingListR: ReadingListR,
	}
}

func (rlu ReadingListUsecase) GetAll(id_user, name string, offset, limit int) ([]response.ReadingList, int, error) {
	readingLists, totalData, err := rlu.ReadingListR.GetAll(id_user, name, offset, limit)

	if err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)
	return readingLists, totalPages, nil
}

func (rlu ReadingListUsecase) GetById(id, user_id string) (*response.ReadingList, error) {
	forum, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (rlu ReadingListUsecase) Create(readingList *entity.ReadingList) error {
	err := rlu.ReadingListR.Create(readingList)
	if err != nil {
		return err
	}
	return nil
}

func (rlu ReadingListUsecase) Update(id, user_id string, readingListId *entity.ReadingList) error {
	_, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return err
	}

	err2 := rlu.ReadingListR.Update(id, user_id, readingListId)

	if err2 != nil {
		return err2
	}
	return nil
}

func (rlu ReadingListUsecase) Delete(id, user_id string) error {
	_, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return err
	}

	err2 := rlu.ReadingListR.Delete(id, user_id)

	if err2 != nil {
		return err2
	}
	return nil
}
