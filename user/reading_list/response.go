package entity

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ReadingList struct {
	ID                  string                      `gorm:"primarykey" json:"id"`
	UserId              string                      `json:"user_id" form:"user_id"`
	Name                string                      `json:"name" form:"name"`
	Description         string                      `json:"description" form:"description"`
	ArticleTotal        int                         `json:"article_total"`
	CreatedAt           time.Time                   `json:"created_at"`
	UpdatedAt           time.Time                   `json:"updated_at"`
	DeletedAt           gorm.DeletedAt              `gorm:"index" json:"deleted_at"`
	ReadingListArticles []entity.ReadingListArticle `gorm:"foreignKey:ReadingListId" json:"reading_list_articles"`
}
