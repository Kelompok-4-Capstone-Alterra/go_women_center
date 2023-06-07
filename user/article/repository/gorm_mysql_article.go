package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetAll(search string, offset, limit int) ([]article.GetAllResponse, int64, error)
	GetById(id string) (article.GetByResponse, error)
	Count() (int, error)
}

type mysqlArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) ArticleRepository {
	return &mysqlArticleRepository{DB: db}
}

func (r *mysqlArticleRepository) GetAll(search string, offset, limit int) ([]article.GetAllResponse, int64, error) {
	var article []article.GetAllResponse
	var count int64
	err := r.DB.Model(&entity.Article{}).
		Where("topic LIKE ? OR title LIKE ? OR author LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&article).Error

	if err != nil {
		return nil, 0, err
	}
	return article, count, nil
}

func (r *mysqlArticleRepository) GetById(id string) (article.GetByResponse, error) {
	var article article.GetByResponse
	err := r.DB.Model(&entity.Article{}).First(&article, "id = ?", id).Error
	if err != nil {
		return article, err
	}

	article.ViewCount++

	err = r.DB.Model(&entity.Article{}).Where("id = ?", id).Updates(entity.Article{ViewCount: article.ViewCount}).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *mysqlArticleRepository) Count() (int, error) {
	var count int64
	err := r.DB.Model(&entity.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
