package counselor

type GetAllResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Topic    string `json:"topic"`
}

type GetByResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Topic          string  `json:"topic"`
	Tarif          float64 `json:"tarif"`
	Description    string  `json:"description"`
}
