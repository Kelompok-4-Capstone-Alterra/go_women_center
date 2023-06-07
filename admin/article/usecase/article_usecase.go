package usecase

import (
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ArticleUsecase interface {
	GetAll(search string, offset, limit int) ([]article.GetAllResponse, int, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (article.GetByResponse, error)
	Create(inputDetail article.CreateRequest, inputImage *multipart.FileHeader) error
	Update(inputDetail article.UpdateRequest, inputImage *multipart.FileHeader) error
	Delete(id string) error
}

type articleUsecase struct {
	articleRepo repository.ArticleRepository
	image       helper.Image
}

func NewArticleUsecase(CRepo repository.ArticleRepository, Image helper.Image) ArticleUsecase {
	return &articleUsecase{articleRepo: CRepo, image: Image}
}

func (u *articleUsecase) GetAll(search string, offset, limit int) ([]article.GetAllResponse, int, error) {

	articles, totalData, err := u.articleRepo.GetAll(search, offset, limit)

	if err != nil {
		return nil, 0, article.ErrInternalServerError
	}

	return articles, helper.GetTotalPages(int(totalData), limit), nil
}

func (u *articleUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.articleRepo.Count()
	if err != nil {
		return 0, article.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func (u *articleUsecase) GetById(id string) (article.GetByResponse, error) {

	articleData, err := u.articleRepo.GetById(id)

	if err != nil {
		return articleData, article.ErrArticleNotFound
	}

	return articleData, nil
}

func (u *articleUsecase) Create(inputDetail article.CreateRequest, inputImage *multipart.FileHeader) error {
	path, err := u.image.UploadImageToS3(inputImage)

	if err != nil {
		return article.ErrInternalServerError
	}
	uuid, _ := helper.NewGoogleUUID().GenerateUUID()

	newArticle := entity.Article{
		ID:          uuid,
		Title:       inputDetail.Title,
		Topic:       constant.TOPICS[inputDetail.Topic],
		Author:      inputDetail.Author,
		Description: inputDetail.Description,
		Image:       path,
	}

	err = u.articleRepo.Create(newArticle)

	if err != nil {
		return article.ErrInternalServerError
	}

	return nil
}

func (u *articleUsecase) Update(inputDetail article.UpdateRequest, inputImage *multipart.FileHeader) error {

	articleData, err := u.articleRepo.GetById(inputDetail.ID)

	if err != nil {
		return article.ErrArticleNotFound
	}

	articleUpdate := entity.Article{
		Title:       inputDetail.Title,
		Topic:       constant.TOPICS[inputDetail.Topic],
		Author:      inputDetail.Author,
		Description: inputDetail.Description,
	}

	if inputImage != nil {
		err := u.image.DeleteImageFromS3(articleData.Image)

		if err != nil {
			return article.ErrInternalServerError
		}

		path, err := u.image.UploadImageToS3(inputImage)

		if err != nil {
			return article.ErrInternalServerError
		}

		articleUpdate.Image = path

	}

	err = u.articleRepo.Update(articleData.ID, articleUpdate)

	if err != nil {
		return article.ErrInternalServerError
	}

	return nil
}

func (u *articleUsecase) Delete(id string) error {

	articleData, err := u.articleRepo.GetById(id)

	if err != nil {
		return article.ErrArticleNotFound
	}

	err = u.articleRepo.Delete(articleData.ID)

	if err != nil {
		return article.ErrInternalServerError
	}

	return nil
}
