package profile

import (
	"mime/multipart"
)

type IdRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

type UpdateRequest struct {
	ID string
	ProfilePicture *multipart.FileHeader `form:"profile_picture" validate:"omitempty"`
	Username       string                `form:"username"`
	Name           string                `form:"name"`
	Email          string                `form:"email" validate:"omitempty,email"`
	PhoneNumber    string                `form:"phone_number"`
	BirthDate      string                `form:"birth_date" `
}

type UpdatePasswordRequest struct {
	ID string
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}