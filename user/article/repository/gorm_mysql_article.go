package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetAll(search, sortBy string) ([]entity.Article, error)
	GetById(id string) (entity.Article, error)
	Count() (int, error)
	UpdateCount(id string, article entity.Article) error
	GetReadingListArticles(id string) ([]entity.ReadingListArticle, error)
}

type mysqlArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) ArticleRepository {
	return &mysqlArticleRepository{DB: db}
}

func (r *mysqlArticleRepository) GetAll(search, sortBy string) ([]entity.Article, error) {
	var articles []entity.Article
	err := r.DB.Model(&entity.Article{}).
		Where("topic LIKE ? OR title LIKE ? OR author LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Order(sortBy).
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *mysqlArticleRepository) GetReadingListArticles(id string) ([]entity.ReadingListArticle, error) {
	var readingListArticles []entity.ReadingListArticle
	err := r.DB.Model(&entity.ReadingListArticle{}).
		Where("user_id LIKE ?", id).
		Find(&readingListArticles).Error

	if err != nil {
		return nil, err
	}
	return readingListArticles, nil
}

func (r *mysqlArticleRepository) GetById(id string) (entity.Article, error) {
	var article entity.Article
	err := r.DB.Model(&entity.Article{}).First(&article, "id = ?", id).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *mysqlArticleRepository) UpdateCount(id string, article entity.Article) error {
	err := r.DB.Model(&entity.Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlArticleRepository) Count() (int, error) {
	var count int64
	err := r.DB.Model(&entity.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
