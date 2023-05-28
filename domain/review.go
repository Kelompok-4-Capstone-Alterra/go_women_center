package domain

type Review struct {
	ID          string  `json:"id" gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID string  `json:"counselor_id" gorm:"type:varchar(36);not null"`
	UserID      string  `json:"user_id" gorm:"type:varchar(36);not null"`
	Rating      float32 `json:"rating" gorm:"type:decimal(2,1)"`
	Comment     string  `json:"comment" gorm:"type:varchar(255)"`
}

type ReviewRepository interface {
	GetAll(idCounselor string, offset, limit int) ([]Review, error)
	Count(idCounselor string) (int, error)
	GetById(id string) (Review, error)
	Save(review Review) error
	GetAverageRating(idCounselor string) (float32, error)
	GetByUserIdAndCounselorId(userId, counselorId string) (Review, error)
}