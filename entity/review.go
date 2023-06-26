package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID          string  `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID string  `gorm:"type:varchar(36);not null"`
	UserID      string  `gorm:"type:varchar(36);not null"`
	Rating      float32 `gorm:"type:decimal(2,1)"`
	Review     string  `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}