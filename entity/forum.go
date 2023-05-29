package entity

import (
	"time"

	"github.com/labstack/echo"
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

type ForumHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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
