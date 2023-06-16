package entity

import (
	"time"
)

type ReadingList struct {
	ID                  string `gorm:"primarykey"`
	UserId              string
	Name                string
	Description         string
	ReadingListArticles []ReadingListArticle `gorm:"foreignKey:ReadingListId"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
