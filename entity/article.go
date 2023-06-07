package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
