package readingList

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID     string `json:"id"`
	Image  string `json:"image"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Topic  string `json:"category"`
}

type ReadingListArticle struct {
	ID            string  `gorm:"primarykey" json:"id"`
	ArticleId     string  `json:"-" form:"article_id"`
	ReadingListId string  `gorm:"type:varchar(50);index" json:"-"`
	Articles      Article `gorm:"foreignKey:ArticleId" json:"article"`
}

type ReadingList struct {
	ID                  string               `gorm:"primarykey" json:"id"`
	UserId              string               `json:"user_id"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	ArticleTotal        int                  `json:"article_total"`
	CreatedAt           time.Time            `json:"created_at"`
	DeletedAt           gorm.DeletedAt       `gorm:"index" json:"-"`
	ReadingListArticles []ReadingListArticle `gorm:"foreignKey:ReadingListId" json:"reading_list_articles"`
}
