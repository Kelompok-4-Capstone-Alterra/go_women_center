package entity

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID         string         `gorm:"primarykey" json:"id"`
	UserId     string         `json:"user_id" form:"user_id"`
	CategoryId uint           `json:"category_id" form:"category_id"`
	Link       string         `json:"link" form:"link"`
	Topic      string         `json:"topic" form:"topic"`
	Status     bool           `json:"status" gorm:"-:all"`
	Member     int            `json:"member" gorm:"-:all"`
	UserForums []UserForum    `gorm:"foreignKey:ForumId" json:"user_forums"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
