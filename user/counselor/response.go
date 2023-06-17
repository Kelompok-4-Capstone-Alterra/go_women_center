package counselor

import "time"

type GetAllResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Price          float64 `json:"price"`
	Rating         float32 `json:"rating"`
}

type GetAllReviewResponse struct {
	ID             string  `json:"id"`
	UserProfile    string  `json:"user_profile"`
	Username       string  `json:"username"`
	Rating         float32 `json:"rating"`
	Review         string  `json:"review"`
	CreatedAt      time.Time  `json:"created_at"`
}

type TimeResponse struct {
	ID     string `json:"id"`
	Time   string `json:"time"`
	Status uint8  `json:"status"`
}

type GetByResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Price          float64 `json:"price"`
	Rating         float32 `json:"rating"`
	Description    string  `json:"description"`
}

type GetReviewByCounselorId struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Rating    float32   `json:"rating"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
}