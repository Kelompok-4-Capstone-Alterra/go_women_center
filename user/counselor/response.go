package counselor

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type GetAllResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Tarif          float64 `json:"tarif"`
	Rating         float32 `json:"rating"`
}

type GetByResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Tarif          float64 `json:"tarif"`
	Rating         float32 `json:"rating"`
	Description    string  `json:"description"`
	Reviews        []entity.Review `json:"reviews" gorm:"foreignKey:CounselorID"`
	Times      		 []entity.Time `json:"times" gorm:"foreignKey:DateID"`
}