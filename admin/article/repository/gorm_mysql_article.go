package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetAll(search, sortBy string, offset, limit int) ([]entity.Article, int64, error)
	GetById(id string) (entity.Article, error)
	Create(article entity.Article) error
	Update(id string, article entity.Article) error
	Delete(id string) error
	Count() (int, error)
	UpdateCount(id string, article entity.Article) error
}

type mysqlArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) ArticleRepository {
	return &mysqlArticleRepository{DB: db}
}

func (r *mysqlArticleRepository) GetAll(search, sortBy string, offset, limit int) ([]entity.Article, int64, error) {
	//TODO: ADD SORT OLDEST AND NEWEST
	var article []entity.Article
	var count int64
	err := r.DB.Model(&entity.Article{}).
		Where("topic LIKE ? OR title LIKE ? OR author LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Order(sortBy).
		Find(&article).Error

	if err != nil {
		return nil, 0, err
	}
	return article, count, nil
}

func (r *mysqlArticleRepository) GetById(id string) (entity.Article, error) {
	var article entity.Article
	err := r.DB.Model(&entity.Article{}).First(&article, "id = ?", id).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *mysqlArticleRepository) Create(article entity.Article) error {
	err := r.DB.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlArticleRepository) Update(id string, article entity.Article) error {

	err := r.DB.Model(&entity.Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlArticleRepository) Delete(id string) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Comment{}).Debug().Unscoped().Delete(&entity.Comment{}, "article_id = ?", id).Error

		if err != nil {
			return err
		}

		err = tx.Model(&entity.ReadingListArticle{}).Debug().Unscoped().Delete(&entity.ReadingListArticle{}, "article_id = ?", id).Error

		if err != nil {
			return err
		}

		err = tx.Model(&entity.Article{}).Debug().Unscoped().Delete(&entity.Article{}, "id = ?", id).Error

		if err != nil {
			return err
		}

		return nil
	})

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

func (r *mysqlArticleRepository) UpdateCount(id string, article entity.Article) error {
	err := r.DB.Model(&entity.Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		return err
	}

	return nil
}
