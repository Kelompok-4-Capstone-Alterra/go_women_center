package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	readingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	repositoryReadingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/repository"
	readingListArticle "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/repository"
)

type ReadingListArticleUsecaseInterface interface {
	Create(createRequest *readingListArticle.CreateRequest) error
	Delete(id, user_id string) error
}

type ReadingListArticleUsecase struct {
	ReadingListArticleR repository.ReadingListArticleRepository
	ReadingListR        repositoryReadingList.ReadingListRepository
}

func NewReadingListArticleUsecase(ReadingListArticleR repository.ReadingListArticleRepository, ReadingListR repositoryReadingList.ReadingListRepository) ReadingListArticleUsecaseInterface {
	return &ReadingListArticleUsecase{
		ReadingListArticleR: ReadingListArticleR,
		ReadingListR:        ReadingListR,
	}
}

func (rlau ReadingListArticleUsecase) Create(createRequest *readingListArticle.CreateRequest) error {
	_, err := rlau.ReadingListR.GetById(createRequest.ReadingListId, createRequest.UserId)
	if err != nil {
		return readingList.ErrFailedGetDetailReadingList
	}

	newReadingListArticle := entity.ReadingListArticle{
		ID:            createRequest.ID,
		ArticleId:     createRequest.ArticleId,
		ReadingListId: createRequest.ReadingListId,
		UserId:        createRequest.UserId,
	}

	err = rlau.ReadingListArticleR.Create(&newReadingListArticle)
	if err != nil {
		return readingListArticle.ErrFailedAddReadingListArticle
	}
	return nil
}

func (rlau ReadingListArticleUsecase) Delete(id, user_id string) error {
	_, err := rlau.ReadingListArticleR.GetById(id, user_id)
	if err != nil {
		return readingListArticle.ErrPageNotFound
	}

	err = rlau.ReadingListArticleR.Delete(id, user_id)

	if err != nil {
		return readingListArticle.ErrFailedDeleteReadingListArticle
	}
	return nil
}
