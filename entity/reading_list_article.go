package entity

import (
	"time"
)

type ReadingListArticle struct {
	ID            string    `gorm:"primarykey" json:"id"`
	ArticleId     string    `json:"article_id" form:"article_id"`
	ReadingListId string    `gorm:"type:varchar(50);index" json:"reading_list_id" form:"reading_list_id"`
	UserId        string    `json:"user_id" form:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
