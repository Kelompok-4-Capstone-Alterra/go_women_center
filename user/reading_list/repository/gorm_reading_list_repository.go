package repository

import (
	"errors"
	"fmt"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	response "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	"gorm.io/gorm"
)

type ReadingListRepository interface {
	GetAll(id_user, name string, offset, limit int) ([]response.ReadingList, int64, error)
	GetById(id, user_id string) (*response.ReadingList, error)
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

func (rlr mysqlReadingListRepository) GetAll(id_user, name string, offset, limit int) ([]response.ReadingList, int64, error) {
	var readingList []response.ReadingList
	var totalData int64 = 3

	err := rlr.DB.Table("reading_lists").Select("reading_lists.id, reading_lists.user_id, reading_lists.name, reading_lists.description, COUNT(reading_list_articles.id) AS article_total").
		Joins("LEFT JOIN reading_list_articles ON reading_lists.id = reading_list_articles.reading_list_id").
		Joins("LEFT JOIN articles ON articles.id = reading_list_articles.article_id").Where("name LIKE ?", "%"+name+"%").
		Group("reading_lists.id").Count(&totalData).Offset(offset).Limit(limit).Preload("ReadingListArticles").Find(&readingList).Error

	if err != nil {
		return nil, totalData, err
	}

	return readingList, totalData, nil
}

func (rlr mysqlReadingListRepository) GetById(id, user_id string) (*response.ReadingList, error) {
	fmt.Println("get by id")
	var readingList response.ReadingList
	err := rlr.DB.Table("reading_lists").Select("reading_lists.id, reading_lists.user_id, reading_lists.name, reading_lists.description, COUNT(reading_list_articles.id) AS article_total").
		Joins("INNER JOIN reading_list_articles ON reading_lists.id = reading_list_articles.reading_list_id").
		Joins("INNER JOIN articles ON articles.id = reading_list_articles.article_id").Where("reading_lists.id = ?", id).
		Preload("ReadingListArticles").First(&readingList).Error

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
	err := rlr.DB.Where("id = ?", id).Take(&entity.ReadingList{}).Error
	if err != nil {
		return err
	}

	err2 := rlr.DB.Model(&readingList).Where("id = ? AND user_id = ? ", id, user_id).Updates(&readingListId).RowsAffected
	if err2 != 1 {
		return errors.New("errors")
	}

	return nil
}

func (rlr mysqlReadingListRepository) Delete(id, user_id string) error {
	err := rlr.DB.Where("id = ?", id).Take(&entity.ReadingList{}).Error

	if err != nil {
		return err
	}

	err2 := rlr.DB.Where("id = ? AND user_id = ? ", id, user_id).Delete(&entity.ReadingList{}).RowsAffected
	if err2 != 1 {
		return errors.New("errors")
	}
	return nil
}
