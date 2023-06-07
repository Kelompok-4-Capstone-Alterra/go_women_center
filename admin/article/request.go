package article

import "mime/multipart"

type CreateRequest struct {
	Title       string                `form:"title" validate:"required"`
	Author      string                `form:"author" validate:"required"`
	Topic       int                   `form:"topic" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description string                `form:"description" validate:"required"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateRequest struct {
	ID          string                `param:"id" validate:"required,uuid"`
	Title       string                `form:"title" validate:"omitempty"`
	Author      string                `form:"author" validate:"omitempty"`
	Topic       int                   `form:"topic" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description string                `form:"description" validate:"omitempty"`
	Image       *multipart.FileHeader `form:"image" validate:"omitempty"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
