package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserReplika struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	Forums    []Forum `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
