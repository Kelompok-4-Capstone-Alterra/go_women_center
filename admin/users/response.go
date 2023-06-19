package users

type GetAllResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetByIdResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	PhoneNumber    string `json:"phone_number"`
}