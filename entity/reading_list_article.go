package entity

import (
	"time"
)

type ReadingListArticle struct {
	ID            string `gorm:"primarykey"`
	ArticleId     string
	ReadingListId string `gorm:"type:varchar(50);index"`
	UserId        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
