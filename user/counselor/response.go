package counselor

type GetAllResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Price          float64 `json:"price"`
	Rating         float32 `json:"rating"`
}

type ReviewResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Username       string  `json:"username"`
	Rating         float32 `json:"rating"`
	Review         string  `json:"review"`
	CreatedAt      string  `json:"created_at"`
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