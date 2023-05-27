package domain

import (
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor"
)

type Counselor struct {
	ID             string  `json:"id,omitempty" form:"id,omitempty" gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	ProfilePicture string  `json:"profile_picture,omitempty" form:"profile_picture,omitempty" gorm:"type:varchar(255)"`
	Username       string  `json:"username,omitempty" form:"username,omitempty" gorm:"type:varchar(150);not null"`
	FullName       string  `json:"full_name,omitempty" form:"full_name,omitempty" gorm:"type:varchar(150);not null"`
	Email          string  `json:"email,omitempty" form:"email,omitempty" gorm:"type:varchar(150);uniqueindex;not null"`
	Topic          string  `json:"topic,omitempty" form:"topic,omitempty" gorm:"type:varchar(50)"`
	Tarif          float64 `json:"tarif,omitempty" form:"tarif,omitempty" gorm:"type:float"`
	Rating         float32 `json:"rating,omitempty" form:"rating,omitempty" gorm:"type:decimal(2,1)"`
	Description    string  `json:"description,omitempty" form:"description,omitempty"`
}

type CounselorUsecase interface {
	GetAll(offset, limit int) ([]Counselor, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (Counselor, error)
	Create(inputDetail counselor.CreateRequest, inputProfilePicture *multipart.FileHeader) error
	Update(inputDetail counselor.UpdateRequest, inputProfilePicture *multipart.FileHeader) error
	Delete(id string) error
}

type CounselorRepository interface {
	GetAll(offset, limit int) ([]Counselor, error)
	Count() (int, error)
	GetByEmail(email string) (Counselor, error)
	GetById(id string) (Counselor, error)
	Create(counselor Counselor) error
	Update(id string, counselor Counselor) error
	Delete(id string) error
}
