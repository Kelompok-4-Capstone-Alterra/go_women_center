package user

type RegisterUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

type VerifyEmailDTO struct {
	Email string `json:"email"`
}
