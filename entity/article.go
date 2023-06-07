package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID           string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	Title        string `gorm:"type:varchar(150);not null"`
	Image        string `gorm:"type:varchar(255)"`
	Author       string `gorm:"type:varchar(150);not null"`
	Topic        string `gorm:"type:varchar(50)"`
	Description  string
	ViewCount    int       `gorm:"type:int"`
	CommentCount int       `gorm:"type:int"`
	Date         time.Time `gorm:"type:date"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
