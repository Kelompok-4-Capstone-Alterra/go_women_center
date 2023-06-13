package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetByArticleId(articleId string, offset, limit int) ([]entity.Comment, int64, error)
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

func (r *mysqlArticleRepository) GetByArticleId(articleId string, offset, limit int) ([]entity.Comment, int64, error) {
	var comments []entity.Comment
	var totalData int64
	err := r.DB.
		Model(&entity.Comment{}).
		Where("article_id = ?", articleId).
		Count(&totalData).
		Offset(offset).
		Limit(limit).
		Find(&comments).Error

	if err != nil {
		return nil, totalData, err
	}

	return comments, totalData, nil
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

	err := r.DB.Delete(&entity.Comment{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
