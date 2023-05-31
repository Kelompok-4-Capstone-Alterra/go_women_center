package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint    `gorm:"primarykey"`
	Name      string  `json:"name" form:"name"`
	Forums    []Forum `gorm:"foreignKey:CategoryId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
