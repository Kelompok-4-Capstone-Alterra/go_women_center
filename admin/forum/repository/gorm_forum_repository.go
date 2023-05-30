package repository

import (
	"fmt"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type mysqlForumRepository struct {
	DB *gorm.DB
}

func NewMysqlForumRepository(db *gorm.DB) entity.ForumRepository {
	return &mysqlForumRepository{DB: db}
}

func (fr mysqlForumRepository) GetAll() ([]entity.Forum, error) {
	var forums []entity.Forum
	err := fr.DB.Find(&forums).Error

	if err != nil {
		return nil, err
	}
	return forums, nil
}

func (fr mysqlForumRepository) GetById(id string) (*entity.Forum, error) {
	var forums entity.Forum
	err := fr.DB.First(&forums, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &forums, nil
}

func (fr mysqlForumRepository) Create(forum *entity.Forum) (*entity.Forum, error) {
	err := fr.DB.Save(forum).Error

	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (fr mysqlForumRepository) Update(id string, forumId *entity.Forum) (*entity.Forum, error) {
	var forum entity.Forum
	fmt.Println("id :", id)
	fmt.Println(forumId)
	err := fr.DB.Model(&forum).Where("id = ?", id).Updates(&forumId).Error
	if err != nil {
		return nil, err
	}
	return forumId, nil
}

func (fr mysqlForumRepository) Delete(id string) error {
	err := fr.DB.Where("id = ?", id).Take(&entity.Forum{}).Error

	if err != nil {
		return err
	}

	err2 := fr.DB.Delete(&entity.Forum{}, &id).Error
	if err != nil {
		return err2
	}
	return nil
}
