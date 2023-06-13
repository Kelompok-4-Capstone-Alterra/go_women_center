package entity

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID         string `gorm:"primarykey"`
	UserId     string
	Category   string
	Link       string
	Topic      string
	Status     bool        `gorm:"-:all"`
	Member     int         `gorm:"-:all"`
	UserForums []UserForum `gorm:"foreignKey:ForumId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
