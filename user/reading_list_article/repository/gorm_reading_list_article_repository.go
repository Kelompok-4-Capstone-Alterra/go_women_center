package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ReadingListArticleRepository interface {
	GetById(id, user_id string) (*entity.ReadingListArticle, error)
	Create(readingListArticle *entity.ReadingListArticle) error
	Delete(id, user_id string) error
}

type mysqlReadingListArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlReadingListArticleRepository(db *gorm.DB) ReadingListArticleRepository {
	return &mysqlReadingListArticleRepository{DB: db}
}

func (rlar mysqlReadingListArticleRepository) GetById(id, user_id string) (*entity.ReadingListArticle, error) {
	var readingListArticle entity.ReadingListArticle
	err := rlar.DB.Where("id = ? AND user_id = ? ", id, user_id).First(&readingListArticle).Error

	if err != nil {
		return nil, err
	}

	return &readingListArticle, nil
}

func (rlar mysqlReadingListArticleRepository) Create(readingListArticle *entity.ReadingListArticle) error {
	err := rlar.DB.Save(readingListArticle).Error

	if err != nil {
		return err
	}
	return nil
}

func (rlar mysqlReadingListArticleRepository) Delete(id, user_id string) error {
	err := rlar.DB.Unscoped().Where("id = ? AND user_id = ? ", id, user_id).Delete(&entity.ReadingListArticle{}).Error
	if err != nil {
		return err
	}
	return nil
}
