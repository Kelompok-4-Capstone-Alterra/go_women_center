package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	ArticleID string `gorm:"type:varchar(36);not null"`
	UserID    string `gorm:"type:varchar(36);not null"`
	Comment   string
	CreatedAt time.Time `gorm:"format:2006-01-02"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
