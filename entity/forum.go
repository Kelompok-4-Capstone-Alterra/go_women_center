package entity

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID         string `gorm:"primarykey"`
	UserId     uint   `json:"user_id" form:"user_id"`
	CategoryId uint   `json:"category_id" form:"category_id"`
	Link       string `json:"link" form:"link"`
	Topic      string `json:"topic" form:"topic"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
