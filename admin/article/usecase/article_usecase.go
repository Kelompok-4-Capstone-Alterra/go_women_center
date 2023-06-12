package usecase

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/repository"
	Comment "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/comment/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	User "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	"golang.org/x/sync/errgroup"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ArticleUsecase interface {
	GetAll(search string, offset, limit int) ([]article.GetAllResponse, int, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (article.GetByResponse, error)
	GetAllComment(id string, offset, limit int) ([]article.CommentResponse, int, error)
	Create(inputDetail article.CreateRequest, inputImage *multipart.FileHeader) error
	Update(inputDetail article.UpdateRequest, inputImage *multipart.FileHeader) error
	Delete(id string) error
	DeleteComment(articleId, commentId string) error
}

type articleUsecase struct {
	articleRepo repository.ArticleRepository
	commentRepo Comment.CommentRepository
	userRepo    User.UserRepository
	image       helper.Image
}

func NewArticleUsecase(ARepo repository.ArticleRepository, CommentRepo Comment.CommentRepository, UserRepo User.UserRepository, Image helper.Image) ArticleUsecase {
	return &articleUsecase{articleRepo: ARepo, commentRepo: CommentRepo, userRepo: UserRepo, image: Image}
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

	dateStr := articleData.Date.Format("2006-01-02")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return articleData, err
	}

	articleDataResponse := article.GetByResponse{
		ID:           articleData.ID,
		Image:        articleData.Image,
		Author:       articleData.Author,
		Topic:        articleData.Topic,
		ViewCount:    articleData.ViewCount,
		CommentCount: articleData.CommentCount,
		Description:  articleData.Description,
		Date:         date,
	}

	return articleDataResponse, nil
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
		Date:        time.Now(),
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

func (u *articleUsecase) GetAllComment(id string, offset, limit int) ([]article.CommentResponse, int, error) {

	// Check article
	_, err := u.articleRepo.GetById(id)
	if err != nil {
		return []article.CommentResponse{}, 0, article.ErrArticleNotFound
	}

	comments, totalData, err := u.commentRepo.GetByArticleId(id, offset, limit)
	if err != nil {
		log.Print(err.Error())
		return []article.CommentResponse{}, 0, article.ErrInternalServerError
	}

	var commentsResponse = make([]article.CommentResponse, len(comments))
	var g errgroup.Group

	for i, comment := range comments {
		i := i
		comment := comment
		g.Go(func() error {
			user, err := u.userRepo.GetById(comment.UserID)
			if err != nil {
				return article.ErrInternalServerError
			}
			commentResponse := article.CommentResponse{
				ID:             comment.ID,
				ArticleID:      comment.ArticleID,
				UserID:         user.ID,
				ProfilePicture: user.ProfilePicture,
				Username:       user.Username,
				Comment:        comment.Comment,
				CreatedAt:      comment.CreatedAt.Format("2006-01-02"),
			}
			commentsResponse[i] = commentResponse

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return []article.CommentResponse{}, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return commentsResponse, totalPages, nil
}

func (u *articleUsecase) DeleteComment(articleId, commentId string) error {
	comment, err := u.commentRepo.GetByArticleIdAndCommentId(articleId, commentId)

	if err != nil {
		return article.ErrCommentNotFound
	}

	articles, err := u.articleRepo.GetById(articleId)
	if err != nil {
		return article.ErrArticleNotFound
	}
	articles.CommentCount--

	commentCount := entity.Article{
		CommentCount: articles.CommentCount,
	}
	u.articleRepo.UpdateCount(articles.ID, commentCount)

	err = u.commentRepo.Delete(comment.ID)
	if err != nil {
		return article.ErrInternalServerError
	}

	return nil
}
