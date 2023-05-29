package entity

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID            uint   `gorm:"primarykey"`
	TopicCategory string `json:"Topic_category" form:"Topic_category"`
	Link          string `json:"link" form:"link"`
	Topic         string `json:"topic" form:"topic"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type ForumUsecase interface {
	GetAll() ([]Forum, error)
	GetById(id string) (Forum, error)
	Create(forum Forum) (Forum, error)
	Update(id string, forumId Forum) (Forum, error)
	Delete(id string) error
}

type ForumRepository interface {
	GetAll() ([]Forum, error)
	GetById(id string) (Forum, error)
	Create(forum Forum) (Forum, error)
	Update(id string, forumId Forum) (Forum, error)
	Delete(id string) error
}
