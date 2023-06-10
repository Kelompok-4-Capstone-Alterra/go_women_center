package entity

import (
	"time"

	"gorm.io/gorm"
)

type ReadingList struct {
	ID          string         `gorm:"primarykey" json:"id"`
	UserId      string         `json:"user_id" form:"user_id"`
	Name        string         `json:"name" form:"name"`
	Description string         `json:"description" form:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
