package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ForumAdminUsecaseInterface interface {
	GetAll(getAllRequest forum.GetAllRequest) ([]forum.ResponseForum, int, error)
	GetById(id string) (*forum.ResponseForum, error)
	Delete(id string) error
}

type ForumAdminUsecase struct {
	ForumAdminR repository.ForumAdminRepository
}

func NewForumAdminUsecase(ForumAdminR repository.ForumAdminRepository) ForumAdminUsecaseInterface {
	return &ForumAdminUsecase{
		ForumAdminR: ForumAdminR,
	}
}

func (fau ForumAdminUsecase) GetAll(getAllRequest forum.GetAllRequest) ([]forum.ResponseForum, int, error) {
	var forums []forum.ResponseForum
	var err error
	var totalData int64

	switch getAllRequest.SortBy {
	case "oldest":
		getAllRequest.SortBy = "forums.created_at ASC"
	case "newest":
		getAllRequest.SortBy = "forums.created_at DESC"
	case "popular":
		getAllRequest.SortBy = "member DESC"
	}

	var newCategory string
	category, ok := constant.TOPICS[getAllRequest.Category]
	if ok {
		newCategory = category[0]
	}

	if getAllRequest.SortBy != "" {
		forums, totalData, err = fau.ForumAdminR.GetAllSortBy(getAllRequest, newCategory)
	} else {
		forums, totalData, err = fau.ForumAdminR.GetAll(getAllRequest, newCategory)
	}

	if err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), getAllRequest.Limit)

	return forums, totalPages, nil
}

func (fau ForumAdminUsecase) GetById(id string) (*forum.ResponseForum, error) {
	forumId, err := fau.ForumAdminR.GetById(id)

	if err != nil {
		return nil, forum.ErrFailedGetDetailReadingList
	} else if forumId.ID == "" {
		return nil, forum.ErrInvalidId
	}
	return forumId, nil
}

func (fau ForumAdminUsecase) Delete(id string) error {
	forumId, err := fau.ForumAdminR.GetById(id)
	if err != nil {
		return err
	} else if forumId.ID == "" {
		return forum.ErrInvalidId
	}

	err = fau.ForumAdminR.Delete(id)

	if err != nil {
		return forum.ErrFailedDeleteReadingList
	}
	return nil
}
