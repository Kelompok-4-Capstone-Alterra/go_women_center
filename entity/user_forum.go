package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserForum struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserId    string         `json:"user_id" form:"user_id"`
	ForumId   string         `gorm:"type:varchar(50);index" json:"forum_id" form:"forum_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
