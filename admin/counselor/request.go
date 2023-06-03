package counselor

import "mime/multipart"

type CreateRequest struct {
	Name           string  `form:"name" validate:"required"`
	Email          string  `form:"email" validate:"required,email"`
	Username       string  `form:"username" validate:"required"`
	Topic          int     `form:"topic" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description    string  `form:"description" validate:"required"`
	Tarif          float64 `form:"tarif" validate:"required,number"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture" validate:"required"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`
	Name        string  `form:"name"`
	Email       string  `form:"email" validate:"omitempty,email"`
	Username    string  `form:"username"`
	Topic       int     `form:"topic" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description string  `form:"description"`
	Tarif       float64 `form:"tarif" validate:"omitempty,number"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}