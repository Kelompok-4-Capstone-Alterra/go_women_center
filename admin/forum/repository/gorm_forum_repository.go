package repository

import (
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
	return []entity.Forum{}, nil
}

func (fr mysqlForumRepository) GetById(id string) (entity.Forum, error) {
	return entity.Forum{}, nil
}

func (fr mysqlForumRepository) Create(forum entity.Forum) (entity.Forum, error) {
	return entity.Forum{}, nil
}

func (fr mysqlForumRepository) Update(id string, forumId entity.Forum) (entity.Forum, error) {
	return entity.Forum{}, nil
}

func (fr mysqlForumRepository) Delete(id string) error {
	return nil
}
