package usecase

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article/repository"
	User "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	Comment "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/comment/repository"
	"golang.org/x/sync/errgroup"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type ArticleUsecase interface {
	GetAll(search string, offset, limit int) ([]article.GetAllResponse, int, error)
	GetById(id string) (article.GetByResponse, error)
	GetAllComment(id string, offset, limit int) ([]article.CommentResponse, int, error)
	CreateComment(input article.CreateCommentRequest) error
	DeleteComment(userId, articleId, commentId string) error
}

type articleUsecase struct {
	articleRepo repository.ArticleRepository
	commentRepo Comment.CommentRepository
	userRepo    User.UserRepository
}

func NewArticleUsecase(ARepo repository.ArticleRepository, CommentRepo Comment.CommentRepository, UserRepo User.UserRepository) ArticleUsecase {
	return &articleUsecase{articleRepo: ARepo, commentRepo: CommentRepo, userRepo: UserRepo}
}

func (u *articleUsecase) GetAll(search string, offset, limit int) ([]article.GetAllResponse, int, error) {

	articles, totalData, err := u.articleRepo.GetAll(search, offset, limit)

	if err != nil {
		return nil, 0, article.ErrInternalServerError
	}

	return articles, helper.GetTotalPages(int(totalData), limit), nil
}

func (u *articleUsecase) GetById(id string) (article.GetByResponse, error) {

	articleData, err := u.articleRepo.GetById(id)

	articleDataResponse := article.GetByResponse{
		ID:           articleData.ID,
		Image:        articleData.Image,
		Author:       articleData.Author,
		Topic:        articleData.Topic,
		ViewCount:    articleData.ViewCount,
		CommentCount: articleData.CommentCount,
		Description:  articleData.Description,
		Date:         articleData.Date.Format("2006-01-01"),
	}

	if err != nil {
		return articleDataResponse, article.ErrArticleNotFound
	}

	articleData.ViewCount++

	viewCount := entity.Article{
		ViewCount: articleData.ViewCount,
	}
	u.articleRepo.UpdateCount(articleData.ID, viewCount)

	return articleDataResponse, nil
}

func (u *articleUsecase) CreateComment(inputComment article.CreateCommentRequest) error {

	articles, err := u.articleRepo.GetById(inputComment.ArticleID)
	if err != nil {
		return article.ErrArticleNotFound
	}
	articles.CommentCount++

	commentCount := entity.Article{
		CommentCount: articles.CommentCount,
	}
	u.articleRepo.UpdateCount(articles.ID, commentCount)

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()

	newComment := entity.Comment{
		ID:        uuid,
		ArticleID: inputComment.ArticleID,
		UserID:    inputComment.UserID,
		Comment:   inputComment.Comment,
	}

	err = u.commentRepo.Save(newComment)

	if err != nil {
		return article.ErrInternalServerError
	}
	return nil
}

func (u *articleUsecase) UpdateCount(inputDetail article.UpdateCountRequest) error {

	articleData, err := u.articleRepo.GetById(inputDetail.ID)

	if err != nil {
		return article.ErrArticleNotFound
	}

	articleUpdate := entity.Article{
		CommentCount: articleData.CommentCount,
		ViewCount:    articleData.ViewCount,
	}
	err = u.articleRepo.UpdateCount(articleData.ID, articleUpdate)

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

func (u *articleUsecase) DeleteComment(userId, articleId, commentId string) error {
	comment, err := u.commentRepo.GetByUserIdAndArticleIdAndCommentId(userId, articleId, commentId)

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
