package article

import "mime/multipart"

type CreateRequest struct {
	
	Topic          int     `form:"topic" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description    string  `form:"description" validate:"required"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture" validate:"required"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`

	Topic       int     `form:"topic" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Description string  `form:"description"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}