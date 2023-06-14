package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	readingListArticle "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/repository"
)

type ReadingListArticleUsecaseInterface interface {
	Create(createRequest []readingListArticle.CreateRequest) error
	Delete(id, user_id string) error
}

type ReadingListArticleUsecase struct {
	ReadingListArticleR repository.ReadingListArticleRepository
}

func NewReadingListArticleUsecase(ReadingListArticleR repository.ReadingListArticleRepository) ReadingListArticleUsecaseInterface {
	return &ReadingListArticleUsecase{
		ReadingListArticleR: ReadingListArticleR,
	}
}

func (rlau ReadingListArticleUsecase) Create(createRequest []readingListArticle.CreateRequest) error {
	var newReadingListArticle []entity.ReadingListArticle
	var InputReadingListArticle entity.ReadingListArticle

	for i := 0; i < len(createRequest); i++ {
		InputReadingListArticle.ID = createRequest[i].ID
		InputReadingListArticle.ArticleId = createRequest[i].ArticleId
		InputReadingListArticle.ReadingListId = createRequest[i].ReadingListId
		InputReadingListArticle.UserId = createRequest[i].UserId
		newReadingListArticle = append(newReadingListArticle, InputReadingListArticle)
	}

	err := rlau.ReadingListArticleR.Create(&newReadingListArticle)
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
