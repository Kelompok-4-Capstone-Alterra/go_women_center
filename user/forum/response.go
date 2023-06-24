package forum

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type User struct {
	ID             string `json:"-"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
}

type ResponseForum struct {
	ID         string             `gorm:"primarykey" json:"id"`
	UserId     string             `json:"user_id" form:"user_id"`
	Users      User               `gorm:"foreignKey:UserId" json:"user"`
	Category   string             `json:"category" form:"category"`
	Link       string             `json:"link" form:"link"`
	Topic      string             `json:"topic" form:"topic"`
	Status     bool               `json:"status"`
	Member     int                `json:"member"`
	CreatedAt  time.Time          `json:"created_at"`
	UserForums []entity.UserForum `gorm:"foreignKey:ForumId" json:"-"`
}
