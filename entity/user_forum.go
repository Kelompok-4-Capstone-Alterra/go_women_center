package entity

import (
	"time"
)

type UserForum struct {
	ID        string    `gorm:"primarykey" json:"id"`
	UserId    string    `json:"user_id" form:"user_id"`
	ForumId   string    `gorm:"type:varchar(50);index" json:"forum_id" form:"forum_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
