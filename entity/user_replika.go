package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserReplika struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Name         string
	Email        string
	Password     string
	PhoneNumber  string
	Birthdate    string
	PhotoProfile string
	Role         string
	Forums       []Forum `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
