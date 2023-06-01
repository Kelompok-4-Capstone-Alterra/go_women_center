package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserForum struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `json:"user_id" form:"user_id"`
	ForumId   string `gorm:"type:varchar(50);index" json:"forum_id" form:"forum_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
