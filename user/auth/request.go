package auth

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	OTP      string `json:"otp" validate:"required,len=4"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyUniqueCredentialRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=4"`
}
