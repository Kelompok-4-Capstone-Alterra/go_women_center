package repository

import (
	"fmt"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	readingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	"gorm.io/gorm"
)

type ReadingListRepository interface {
	GetAll(getAllParams readingList.GetAllRequest) ([]readingList.ReadingList, int64, error)
	GetById(id, user_id string) (*readingList.ReadingList, error)
	Create(readingList *entity.ReadingList) error
	Update(id, user_id string, readingList *entity.ReadingList) error
	Delete(id, user_id string) error
}

type mysqlReadingListRepository struct {
	DB *gorm.DB
}

func NewMysqlReadingListRepository(db *gorm.DB) ReadingListRepository {
	return &mysqlReadingListRepository{DB: db}
}

func (rlr mysqlReadingListRepository) GetAll(getAllParams readingList.GetAllRequest) ([]readingList.ReadingList, int64, error) {
	fmt.Println(getAllParams.Name)
	var dataReadingList []readingList.ReadingList
	var totalData int64

	err := rlr.DB.Table("reading_lists").Select("reading_lists.id, reading_lists.user_id, reading_lists.name, reading_lists.description, COUNT(reading_list_articles.id) AS article_total").
		Joins("LEFT JOIN reading_list_articles ON reading_lists.id = reading_list_articles.reading_list_id").
		Joins("LEFT JOIN articles ON articles.id = reading_list_articles.article_id").Where("reading_lists.user_id = ? AND reading_lists.name LIKE ?", getAllParams.UserId, "%"+getAllParams.Name+"%").
		Group("reading_lists.id").Order(getAllParams.SortBy).Count(&totalData).Offset(getAllParams.Offset).Limit(getAllParams.Limit).Preload("ReadingListArticles.Articles").Find(&dataReadingList).Error

	if err != nil {
		return nil, totalData, err
	}

	return dataReadingList, totalData, nil
}

func (rlr mysqlReadingListRepository) GetById(id, user_id string) (*readingList.ReadingList, error) {

	var readingList readingList.ReadingList
	err := rlr.DB.Table("reading_lists").Select("reading_lists.id, reading_lists.user_id, reading_lists.name, reading_lists.description, COUNT(reading_list_articles.id) AS article_total").
		Joins("LEFT JOIN reading_list_articles ON reading_lists.id = reading_list_articles.reading_list_id").
		Joins("LEFT JOIN articles ON articles.id = reading_list_articles.article_id").Where("reading_lists.user_id = ?", user_id).
		Group("reading_lists.id").
		Preload("ReadingListArticles.Articles").First(&readingList, "reading_lists.id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &readingList, nil
}

func (rlr mysqlReadingListRepository) Create(readingList *entity.ReadingList) error {
	err := rlr.DB.Save(readingList).Error

	if err != nil {
		return err
	}
	return nil
}

func (rlr mysqlReadingListRepository) Update(id, user_id string, readingListId *entity.ReadingList) error {
	var readingList entity.ReadingList
	err := rlr.DB.Model(&readingList).Where("id = ? AND user_id = ? ", id, user_id).Updates(&readingListId).Error

	if err != nil {
		return err
	}

	return nil
}

func (rlr mysqlReadingListRepository) Delete(id, user_id string) error {
	err := rlr.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.ReadingListArticle{}).Unscoped().Delete(&entity.ReadingListArticle{}, "reading_list_id = ?", id).Error

		if err != nil {
			return err
		}
		err = tx.Model(&entity.ReadingList{}).Unscoped().Delete(&entity.ReadingList{}, "id = ?", id).Error

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
