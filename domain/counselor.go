package domain

type Counselor struct {
	ID           string  `json:"id" gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	PhotoProfile string  `json:"photo_profile,omitempty" gorm:"type:varchar(255)"`
	FullName     string  `json:"full_name,omitempty" gorm:"type:varchar(150);not null"`
	Email        string  `json:"email,omitempty" gorm:"type:varchar(150);uniqueindex;not null"`
	Tarif        int     `json:"tarif,omitempty" gorm:"type:float64"`
	Rating       float32 `json:"rating,omitempty" gorm:"type:decimal(2,1)"`
	Category     string  `json:"category,omitempty" gorm:"type:varchar(150)"`
	Description  string  `json:"description,omitempty"`
	Topic        string  `json:"topic,omitempty" gorm:"type:varchar(150)"`
}

type CounselorUsecase interface {
	GetAll(page, limit int) ([]Counselor, error)
	GetTotalPages(limit int) (int, error)
	// GetById(id string) (Counselor, error)
	Create(counselor Counselor) error
	// Update(id string, counselor Counselor) error
	// Delete(id string) error
}

type CounselorRepository interface {
	GetAll(offset, limit int) ([]Counselor, error)
	Count() (int, error)
	// GetById(id string) (Counselor, error)
	Create(counselor Counselor) error
	// Update(id string, counselor Counselor) error
	// Delete(id string) error
}
