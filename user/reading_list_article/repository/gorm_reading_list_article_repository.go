package repository

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ReadingListArticleRepository interface {
	Create(readingListArticle *entity.ReadingListArticle) error
	Delete(id, user_id string) error
}

type mysqlReadingListArticleRepository struct {
	DB *gorm.DB
}

func NewMysqlReadingListArticleRepository(db *gorm.DB) ReadingListArticleRepository {
	return &mysqlReadingListArticleRepository{DB: db}
}

func (rlar mysqlReadingListArticleRepository) Create(readingListArticle *entity.ReadingListArticle) error {
	err := rlar.DB.Save(readingListArticle).Error

	if err != nil {
		return err
	}
	return nil
}

func (rlar mysqlReadingListArticleRepository) Delete(id, user_id string) error {
	err := rlar.DB.Where("id = ?", id).Take(&entity.ReadingListArticle{}).Error

	if err != nil {
		return err
	}

	err2 := rlar.DB.Where("id = ? AND user_id = ? ", id, user_id).Delete(&entity.ReadingListArticle{}).RowsAffected
	if err2 != 1 {
		return errors.New("errors")
	}
	return nil
}
