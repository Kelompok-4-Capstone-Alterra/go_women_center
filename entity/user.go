package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string     `gorm:"type:varchar(36);primaryKey;uniqueindex;not null"`
	Username       string     `gorm:"type:varchar(150);uniqueindex;not null"`
	Name           string     `gorm:"type:varchar(150);not null"`
	Email          string     `gorm:"type:varchar(150);uniqueindex;not null"`
	Password       string     `gorm:"type:varchar(64)"`
	PhoneNumber    string     `gorm:"type:varchar(20)"`
	ProfilePicture string     `gorm:"type:varchar(255)"`
	BirthDate      *time.Time `gorm:"type:date"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Reviews        []Review       `gorm:"foreignKey:UserID;references:ID"`
	Forums         []Forum        `gorm:"foreignKey:UserId"`
	UserForums     []UserForum    `gorm:"foreignKey:UserId"`
}

type OTP struct {
	Token    string
	Deadline time.Time
}

func NewOTP(token string) OTP {
	return OTP{
		Token:    token,
		Deadline: time.Now().Add(2 * time.Minute),
	}
}
