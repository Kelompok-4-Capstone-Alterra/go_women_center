package entity

import (
	"time"

	"gorm.io/gorm"
)

type Career struct {
	ID            string  `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	Image         string  `gorm:"type:varchar(255)"`
	JobPosition   string  `gorm:"type:varchar(150);not null"`
	CompanyName   string  `gorm:"type:varchar(150);not null"`
	Location      string  `gorm:"type:varchar(150);not null"`
	MinSalary     float64 `gorm:"type:float"`
	MaxSalary     float64 `gorm:"type:float"`
	MinExperience string  `gorm:"type:varchar(150);not null"`
	LastEducation string  `gorm:"type:varchar(150);not null"`
	CompanyEmail  string  `gorm:"type:varchar(150);not null"`
	Description   string
	Requirement   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
