package entity

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
)

type Transaction struct {
	// TODO: string length fix
	ID                 string `gorm:"primarykey"`
	UserId             string `gorm:"type:varchar(36);not null"`
	CounselorId        string `gorm:"type:varchar(36);not null"`
	DateId             string `gorm:"type:varchar(36);not null"`
	TimeId             string `gorm:"type:varchar(36);not null"`
	CounselorTopic     string `gorm:"type:varchar(50)"`
	Link               string // link meeting
	TimeStart          string // Convert from time to valid string first
	ConsultationMethod string
	Status             transaction.Status `gorm:"type:varchar(36)"`
	ValueVoucher       int64
	GrossPrice         int64
	TotalPrice         int64
	IsReviewed         bool
	Created_at         time.Time
}
