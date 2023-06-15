package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	readingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/repository"
)

type ReadingListUsecaseInterface interface {
	GetAll(getAllParams readingList.GetAllRequest) ([]readingList.ReadingList, int, error)
	GetById(id, user_id string) (*readingList.ReadingList, error)
	Create(createForum *readingList.CreateRequest) error
	Update(id, user_id string, readingList *readingList.UpdateRequest) error
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

func (rlu ReadingListUsecase) GetAll(getAllParams readingList.GetAllRequest) ([]readingList.ReadingList, int, error) {
	switch getAllParams.Sort {
	case "oldest":
		getAllParams.Sort = "reading_lists.created_at ASC"
	case "newest":
		getAllParams.Sort = "reading_lists.created_at DESC"
	default:
		getAllParams.Sort = "reading_lists.created_at"
	}
	readingLists, totalData, err := rlu.ReadingListR.GetAll(getAllParams)

	if err != nil {
		return nil, 0, readingList.ErrFailedGetReadingList
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllParams.Limit)
	return readingLists, totalPages, nil
}

func (rlu ReadingListUsecase) GetById(id, user_id string) (*readingList.ReadingList, error) {
	forum, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return nil, readingList.ErrFailedGetDetailReadingList
	} else if forum.ID == "" {
		return nil, readingList.ErrPageNotFound
	}

	return forum, nil
}

func (rlu ReadingListUsecase) Create(createForum *readingList.CreateRequest) error {
	forum := entity.ReadingList{
		ID:          createForum.ID,
		UserId:      createForum.UserId,
		Name:        createForum.Name,
		Description: createForum.Description,
	}
	err := rlu.ReadingListR.Create(&forum)
	if err != nil {
		return readingList.ErrFailedCreateReadingList
	}

	return nil
}

func (rlu ReadingListUsecase) Update(id, user_id string, readingListId *readingList.UpdateRequest) error {
	forum, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return readingList.ErrFailedGetDetailReadingList
	} else if forum.ID == "" {
		return readingList.ErrPageNotFound
	}

	newReadingList := entity.ReadingList{
		Name:        readingListId.Name,
		Description: readingListId.Description,
	}
	err = rlu.ReadingListR.Update(id, user_id, &newReadingList)

	if err != nil {
		return readingList.ErrFailedUpdateReadingList
	}

	return nil
}

func (rlu ReadingListUsecase) Delete(id, user_id string) error {
	forum, err := rlu.ReadingListR.GetById(id, user_id)

	if err != nil {
		return readingList.ErrFailedGetDetailReadingList
	} else if forum.ID == "" {
		return readingList.ErrPageNotFound
	}

	err = rlu.ReadingListR.Delete(id, user_id)

	if err != nil {
		return err
	}
	return nil
}
