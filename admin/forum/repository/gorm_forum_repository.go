package repository

import (
	"errors"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ForumRepository interface {
	Delete(id string) error
}

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) Delete(id string) error {
	err := fr.DB.Where("id = ?", id).Take(&entity.Forum{}).Error

	if err != nil {
		return err
	}

	err2 := fr.DB.Where("id = ? ", id).Delete(&entity.Forum{}).RowsAffected
	if err2 != 1 {
		return errors.New("errors")
	}
	return nil
}
