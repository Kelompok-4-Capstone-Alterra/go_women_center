package counselor

import "mime/multipart"

type CreateRequest struct {
	Name           string  `form:"name" validate:"required"`
	Email          string  `form:"email" validate:"required,email"`
	Username       string  `form:"username" validate:"required"`
	Topic          int     `form:"topic" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description    string  `form:"description" validate:"required"`
	Price          float64 `form:"price" validate:"required,number"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture" validate:"required"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`
	Name        string  `form:"name"`
	Email       string  `form:"email" validate:"omitempty,email"`
	Username    string  `form:"username"`
	Topic       int     `form:"topic" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description string  `form:"description"`
	Price       float64 `form:"price" validate:"omitempty,number"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture"`
}

type GetAllRequest struct {
	Page int `query:"page" validate:"omitempty"`
	Limit int `query:"limit" validate:"omitempty"`
	Search string `query:"search" validate:"omitempty"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=oldest newest"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}