package entity

import (
	"time"

	"gorm.io/gorm"
)

type Date struct {
	ID          string `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID string `gorm:"type:varchar(36);not null"`
	Date 			 time.Time `gorm:"type:date"`
	Times 		 []Time `gorm:"foreignKey:DateID"`
}

func(s *Date) BeforeDelete(tx *gorm.DB) error {
	tx.Model(&Time{}).Where("date_id = ?", s.ID).Delete(&Time{})
	return nil
}