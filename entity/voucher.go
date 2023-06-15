package entity

import "time"

type Voucher struct {
	ID        string    `gorm:"primarykey" json:"id"`
	UserId    string    `json:"user_id" form:"user_id"`
	ExpDate   time.Time `json:"exp_date"`
	Value     int64     `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
