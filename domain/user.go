package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID string `gorm:"primaryKey"` 
	Username string
	Name string
	Email string
	Password string
	PhoneNumber string
	Birthdate string
	PhotoProfile string
	Role string
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OTP struct {
	Token string
	Deadline time.Time
}

func NewOTP(token string) OTP {
	return OTP{
		Token: token,
		Deadline: time.Now().Add(2 * time.Minute),
	}
}