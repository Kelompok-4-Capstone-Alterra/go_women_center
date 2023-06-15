package entity

import (
	"time"
)

type Transaction struct {
	// TODO: string length fix
	ID                 string `gorm:"primarykey"`
	UserId             string `gorm:"type:varchar(36);not null"`
	CounselorID        string `gorm:"type:varchar(36);not null"`
	DateId             string `gorm:"type:varchar(36);not null"`
	TimeId             string `gorm:"type:varchar(36);not null"`
	CounselorTopic     string `gorm:"type:varchar(50)"`
	Link               string // link meeting
	TimeStart          string `gorm:"type:time(0);not null"`
	ConsultationMethod string
	Status             string `gorm:"type:varchar(36)"`
	ValueVoucher       int64
	GrossPrice         int64
	TotalPrice         int64
	IsReviewed         bool
	Created_at         time.Time
}
