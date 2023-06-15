package entity

import (
	"time"
)

type Transaction struct {
	// TODO: string length fix
	ID                 string    `json:"id" gorm:"primarykey"`
	UserId             string    `json:"user_id" gorm:"type:varchar(36);not null"`
	CounselorID        string    `json:"counselor_id" gorm:"type:varchar(36);not null"`
	DateId             string    `json:"date_id" gorm:"type:varchar(36);not null"`
	TimeId             string    `json:"time_id" gorm:"type:varchar(36);not null"`
	CounselorTopic     string    `json:"counselor_topic" gorm:"type:varchar(50)"`
	Link               string    `json:"link" ` // link meeting
	TimeStart          string    `json:"time_start" gorm:"type:time(0);not null"`
	ConsultationMethod string    `json:"consultation_method" `
	Status             string    `json:"status" gorm:"type:varchar(36)"`
	ValueVoucher       int64     `json:"value_voucher" `
	GrossPrice         int64     `json:"gross_price" `
	TotalPrice         int64     `json:"total_price" `
	IsReviewed         bool      `json:"is_reviewed" `
	Created_at         time.Time `json:"created_at" `
}
