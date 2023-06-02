package review

import "time"

type GetAllResponse struct {
	ID          string  `json:"id"`
	CounselorID string  `json:"counselor_id"`
	UserID      string  `json:"user_id"`
	Rating      float32 `json:"rating"`
	Review     string  `json:"review"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetByCounselorId struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Rating      float32 `json:"rating"`
	Review     string  `json:"review"`
	CreatedAt   time.Time `json:"created_at"`
}