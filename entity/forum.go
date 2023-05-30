package entity

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID            uint   `gorm:"primarykey"`
	UserId        uint   `json:"user_id" form:"user_id"`
	TopicCategory string `json:"Topic_category" form:"Topic_category"`
	Link          string `json:"link" form:"link"`
	Topic         string `json:"topic" form:"topic"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
