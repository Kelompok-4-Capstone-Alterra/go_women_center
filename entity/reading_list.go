package entity

import (
	"time"

	"gorm.io/gorm"
)

type ReadingList struct {
	ID                  string `gorm:"primarykey"`
	UserId              string
	Name                string
	Description         string
	ReadingListArticles []ReadingListArticle `gorm:"foreignKey:ReadingListId"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
