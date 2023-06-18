package profile

type GetByIdResponse struct {
	ID             string `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
}
