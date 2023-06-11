package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID                  string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	Image               string `gorm:"type:varchar(255)"`
	Title               string `gorm:"type:varchar(150);not null"`
	Author              string `gorm:"type:varchar(150);not null"`
	Topic               string `gorm:"type:varchar(50)"`
	Description         string
	Comments            []Comment            `gorm:"foreignkey:ArticleID"`
	ReadinglistArticles []ReadingListArticle `gorm:"foreignKey:ArticleId" json:"reading_list"`
	Date                time.Time            `gorm:"type:date"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
