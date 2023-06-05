package entity

import "gorm.io/gorm"

type Counselor struct {
	ID             string  `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	ProfilePicture string  `gorm:"type:varchar(255)"`
	Username       string  `gorm:"type:varchar(150);uniqueindex;not null"`
	Name           string  `gorm:"type:varchar(150);not null"`
	Email          string  `gorm:"type:varchar(150);uniqueindex;not null"`
	Topic          string  `gorm:"type:varchar(50)"`
	Price          float64 `gorm:"type:float"`
	Rating         float32 `gorm:"type:decimal(2,1)"`
	Description    string
	Reviews        []Review `gorm:"foreignkey:CounselorID"`
	Dates          []Date `gorm:"foreignkey:CounselorID"`
	Times 		   []Time `gorm:"foreignkey:CounselorID"`
}

func(c *Counselor) BeforeDelete(tx *gorm.DB) error {
	tx.Model(&Review{}).Where("counselor_id = ?", c.ID).Delete(&Review{})
	tx.Model(&Date{}).Where("counselor_id = ?", c.ID).Delete(&Date{})
	tx.Model(&Time{}).Where("counselor_id = ?", c.ID).Delete(&Time{})
	return nil
}