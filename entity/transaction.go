package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	// TODO: string length fix
	ID                 string `gorm:"primarykey"`
	UserId             string
	Date               time.Time `gorm:"type:date"`
	CounselorId        string    `gorm:"type:varchar(36);not null"`
	Link               string    // link meeting
	CounselorTopic     string    `gorm:"type:varchar(50)"`
	TimeId             string    `gorm:"type:varchar(36)"`
	TimeStart          string    // Convert from time to valid string first
	ConsultationMethod string
	Status             string // TODO: change to enum gorm
	ValueVoucher       float64
	GrossPrice         float64
	TotalPrice         float64
	IsReviewed         bool
	Created_at         time.Time
	Deleted_at         gorm.DeletedAt
}
