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
	GetAll(search, userId, sortBy string) ([]article.GetAllResponse, error)
	GetById(id string) (article.GetByResponse, error)
	GetAllComment(id string) ([]article.CommentResponse, error)
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

func (u *articleUsecase) GetAll(search, userId, sortBy string) ([]article.GetAllResponse, error) {
	switch sortBy {
	case "newest":
		sortBy = "date DESC"
	case "oldest":
		sortBy = "date ASC"
	case "most_viewed":
		sortBy = "view_count DESC"
	}

	articles, err := u.articleRepo.GetAll(search, sortBy)
	if err != nil {
		log.Print(err.Error())
		return []article.GetAllResponse{}, article.ErrInternalServerError
	}

	readingListArticles, err := u.GetReadingListArticles(userId)

	if err != nil {
		log.Print(err.Error())
		return []article.GetAllResponse{}, article.ErrInternalServerError
	}

	var articlesResponse = make([]article.GetAllResponse, len(articles))
	var g errgroup.Group

	for i, articles := range articles {
		i := i
		articles := articles
		g.Go(func() error {
			articleResponse := article.GetAllResponse{
				ID:           articles.ID,
				Image:        articles.Image,
				Author:       articles.Author,
				Topic:        articles.Topic,
				Title:        articles.Title,
				ViewCount:    articles.ViewCount,
				CommentCount: articles.CommentCount,
				Description:  articles.Description,
				Date:         articles.Date.Format("2006-01-02"),
			}
			articlesResponse[i] = articleResponse

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return []article.GetAllResponse{}, err
	}

	for i, articleData := range articlesResponse {
		for _, ReadingListData := range readingListArticles {
			if articleData.ID == ReadingListData.ArticleID {
				articlesResponse[i].Saved = true
				break
			}
		}
	}

	return articlesResponse, nil
}

func (u *articleUsecase) GetReadingListArticles(id string) ([]article.ReadingListArticleResponse, error) {

	readingListArticles, err := u.articleRepo.GetReadingListArticles(id)
	if err != nil {
		log.Print(err.Error())
		return []article.ReadingListArticleResponse{}, article.ErrInternalServerError
	}

	var ReadingListArticlesResponse = make([]article.ReadingListArticleResponse, len(readingListArticles))
	var g errgroup.Group

	for i, ReadingListArticle := range readingListArticles {
		i := i
		ReadingListArticle := ReadingListArticle
		g.Go(func() error {
			ReadingListArticle := article.ReadingListArticleResponse{
				ID:            ReadingListArticle.ID,
				ArticleID:     ReadingListArticle.ArticleId,
				UserID:        ReadingListArticle.UserId,
				ReadingListID: ReadingListArticle.ReadingListId,
			}
			ReadingListArticlesResponse[i] = ReadingListArticle
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return []article.ReadingListArticleResponse{}, err
	}

	return ReadingListArticlesResponse, nil
}

func (u *articleUsecase) GetById(id string) (article.GetByResponse, error) {

	articleData, err := u.articleRepo.GetById(id)

	articleDataResponse := article.GetByResponse{
		ID:           articleData.ID,
		Image:        articleData.Image,
		Author:       articleData.Author,
		Topic:        articleData.Topic,
		Title:        articleData.Title,
		CommentCount: articleData.CommentCount,
		Description:  articleData.Description,
		Date:         articleData.Date.Format("2006-01-02"),
	}

	if err != nil {
		return articleDataResponse, article.ErrArticleNotFound
	}

	articleData.ViewCount++

	viewCount := entity.Article{
		ViewCount: articleData.ViewCount,
	}
	u.articleRepo.UpdateCount(articleData.ID, viewCount)

	articleDataResponse.ViewCount = articleData.ViewCount

	return articleDataResponse, nil
}

func (u *articleUsecase) CreateComment(inputComment article.CreateCommentRequest) error {

	articles, err := u.articleRepo.GetById(inputComment.ArticleID)
	if err != nil {
		return article.ErrArticleNotFound
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()

	newComment := entity.Comment{
		ID:        uuid,
		ArticleID: inputComment.ArticleID,
		UserID:    inputComment.UserID,
		Comment:   inputComment.Comment,
	}
	articles.CommentCount++

	commentCount := entity.Article{
		CommentCount: articles.CommentCount,
	}
	u.articleRepo.UpdateCount(articles.ID, commentCount)

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

func (u *articleUsecase) GetAllComment(id string) ([]article.CommentResponse, error) {
	// Check article
	_, err := u.articleRepo.GetById(id)
	if err != nil {
		return []article.CommentResponse{}, article.ErrArticleNotFound
	}

	comments, err := u.commentRepo.GetByArticleId(id)
	if err != nil {
		log.Print(err.Error())
		return []article.CommentResponse{}, article.ErrInternalServerError
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
		return []article.CommentResponse{}, err
	}

	return commentsResponse, nil
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
