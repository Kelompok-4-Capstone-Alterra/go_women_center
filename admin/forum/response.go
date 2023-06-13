package forum

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ResponseForum struct {
	ID             string             `gorm:"primarykey" json:"id"`
	Name           string             `json:"name" form:"name"`
	ProfilePicture string             `json:"profile_picture" form:"profile_picture"`
	Category       string             `json:"category" form:"category"`
	Link           string             `json:"link" form:"link"`
	Topic          string             `json:"topic" form:"topic"`
	Member         int                `json:"member"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	UserForums     []entity.UserForum `gorm:"foreignKey:ForumId" json:"user_forums"`
}
