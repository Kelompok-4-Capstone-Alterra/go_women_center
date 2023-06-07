package response

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ResponseForum struct {
	ID         string             `gorm:"primarykey" json:"id"`
	UserId     string             `json:"user_id" form:"user_id"`
	CategoryId uint               `json:"category_id" form:"category_id"`
	Link       string             `json:"link" form:"link"`
	Topic      string             `json:"topic" form:"topic"`
	Status     bool               `json:"status"`
	Member     int                `json:"member"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DeletedAt  gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	UserForums []entity.UserForum `gorm:"foreignKey:ForumId" json:"user_forums"`
}

type ResponseForumDetail struct {
	ID         string             `gorm:"primarykey" json:"id"`
	UserId     string             `json:"user_id" form:"user_id"`
	CategoryId uint               `json:"category_id" form:"category_id"`
	Link       string             `json:"link" form:"link"`
	Topic      string             `json:"topic" form:"topic"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DeletedAt  gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	UserForums []entity.UserForum `gorm:"foreignKey:ForumId" json:"user_forums"`
}
