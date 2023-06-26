package entity

import (
	"time"

	"gorm.io/gorm"
)

type Counselor struct {
	ID             string         `json:"id" gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	ProfilePicture string         `json:"profile_picture" gorm:"type:varchar(255)"`
	Username       string         `json:"username" gorm:"type:varchar(150);index;not null"`
	Name           string         `json:"name" gorm:"type:varchar(150);not null"`
	Email          string         `json:"email" gorm:"type:varchar(150);index;not null"`
	Topic          string         `json:"topic" gorm:"type:varchar(50)"`
	Price          float64        `json:"price" gorm:"type:float"`
	Rating         float32        `json:"rating" gorm:"type:decimal(2,1)"`
	Description    string         `json:"description"`
	Reviews        []Review       `json:"reviews" gorm:"foreignkey:CounselorID"`
	Dates          []Date         `json:"dates" gorm:"foreignkey:CounselorID"`
	Times          []Time         `json:"times" gorm:"foreignkey:CounselorID"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	CreatedAt      time.Time      `json:"created_at"`
	Transactions   []Transaction  `json:"transactions" gorm:"foreignkey:CounselorID"`
}
