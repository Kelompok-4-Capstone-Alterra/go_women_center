package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetByArticleId(articleId string) ([]entity.Comment, error)
	Save(comment entity.Comment) error
	GetByUserIdAndArticleId(userId, articleId string) (entity.Comment, error)
	GetByUserId(userId string) (entity.Comment, error)
	GetByUserIdAndArticleIdAndCommentId(userId, articleId, commentId string) (entity.Comment, error)
	Delete(id string) error
}

type mysqlArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) CommentRepository {
	return &mysqlArticleRepository{DB: db}
}

func (r *mysqlArticleRepository) GetByArticleId(articleId string) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.DB.
		Model(&entity.Comment{}).
		Where("article_id = ?", articleId).
		Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *mysqlArticleRepository) Save(comment entity.Comment) error {
	err := r.DB.Save(&comment).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlArticleRepository) GetByUserIdAndArticleId(userId, articleId string) (entity.Comment, error) {
	comment := entity.Comment{}

	err := r.DB.Where("user_id = ? AND article_id = ?", userId, articleId).First(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *mysqlArticleRepository) GetByUserIdAndArticleIdAndCommentId(userId, articleId, commentId string) (entity.Comment, error) {
	comment := entity.Comment{}

	err := r.DB.Where("user_id = ? AND  article_id= ? AND  id= ?", userId, articleId, commentId).First(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *mysqlArticleRepository) GetByUserId(userId string) (entity.Comment, error) {
	comment := entity.Comment{}

	err := r.DB.Where("user_id = ?", userId).First(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *mysqlArticleRepository) Delete(id string) error {

	err := r.DB.Unscoped().Delete(&entity.Comment{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
