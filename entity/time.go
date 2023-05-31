package entity

import "time"

type Time struct {
	ID   string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	Time time.Time `gorm:"type:time(3);not null"`
	DateID string `gorm:"type:varchar(36);not null"`
}
