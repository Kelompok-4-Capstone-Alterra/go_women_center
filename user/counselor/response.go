package counselor

type GetAllResponse struct {
	ID             string  `json:"id"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Topic          string  `json:"topic"`
	Tarif          float64 `json:"tarif"`
	Rating         float32 `json:"rating"`
}