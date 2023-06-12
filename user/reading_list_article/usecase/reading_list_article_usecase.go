package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/repository"
)

type ReadingListArticleUsecaseInterface interface {
	Create(readingListArticle *entity.ReadingListArticle) error
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

func (rlau ReadingListArticleUsecase) Create(readingListArticle *entity.ReadingListArticle) error {
	err := rlau.ReadingListArticleR.Create(readingListArticle)
	if err != nil {
		return err
	}
	return nil
}

func (rlau ReadingListArticleUsecase) Delete(id, user_id string) error {
	err := rlau.ReadingListArticleR.Delete(id, user_id)

	if err != nil {
		return err
	}
	return nil
}
