package readingList

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	Image       string `gorm:"type:varchar(255)"`
	Title       string `gorm:"type:varchar(150);not null"`
	Author      string `gorm:"type:varchar(150);not null"`
	Topic       string `gorm:"type:varchar(50)"`
	Description string
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ReadingListArticle struct {
	ID            string         `gorm:"primarykey" json:"id"`
	ArticleId     string         `json:"article_id" form:"article_id"`
	ReadingListId string         `gorm:"type:varchar(50);index" json:"reading_list_id" form:"reading_list_id"`
	UserId        string         `json:"user_id" form:"user_id"`
	Articles      Article        `gorm:"foreignKey:ArticleId"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ReadingList struct {
	ID                  string               `gorm:"primarykey" json:"id"`
	UserId              string               `json:"user_id" form:"user_id"`
	Name                string               `json:"name" form:"name"`
	Description         string               `json:"description" form:"description"`
	ArticleTotal        int                  `json:"article_total"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
	ReadingListArticles []ReadingListArticle `gorm:"foreignKey:ReadingListId" json:"reading_list_articles"`
}
