package entity

import (
	"time"
)

type Date struct {
	ID           string        `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID  string        `gorm:"type:varchar(36);not null"`
	Date         time.Time     `gorm:"type:date"`
}