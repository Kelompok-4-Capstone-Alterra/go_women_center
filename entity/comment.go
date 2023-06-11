package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string         `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	ArticleID string         `json:"article_id" form:"article_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
