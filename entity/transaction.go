package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                 string `gorm:"primarykey"`
	UserId             string
	ScheduleId         string
	CounselorId        string
	Link               string
	CounselorTopic     string
	TimeStart          string
	ConsultationMethod string
	Status             string
	ValueVoucher       float64
	GrossPrice         int
	TotalPrice         float64
	Created_at         time.Time
	Deleted_at         gorm.DeletedAt
}
