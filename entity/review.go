package entity

import "gorm.io/gorm"

type Review struct {
	ID          string  `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID string  `gorm:"type:varchar(36);not null"`
	UserID      string  `gorm:"type:varchar(36);not null"`
	Rating      float32 `gorm:"type:decimal(2,1)"`
	Comment     string  `gorm:"type:varchar(255)"`
}

func (r *Review) AfterCreate(tx *gorm.DB) error {
	// Update counselor rating
	var avgRating float32
	err := tx.Model(&Review{}).Where("counselor_id = ?", r.CounselorID).Select("AVG(rating)").Scan(&avgRating).Error
	
	if err != nil {
		return err
	}

	tx.Model(&Counselor{}).Where("id = ?", r.CounselorID).Update("rating", avgRating)
	return nil
}