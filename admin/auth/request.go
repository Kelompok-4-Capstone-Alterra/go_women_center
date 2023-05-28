package auth

type LoginAdminDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}