package career

import "mime/multipart"

type CreateRequest struct {
	JobPosition   string                `form:"job_position" validate:"required"`
	JobType       string                `form:"job_type" validate:"required"`
	CompanyName   string                `form:"company_name" validate:"required"`
	Location      string                `form:"location" validate:"required"`
	Salary        float64               `form:"salary" validate:"omitempty"`
	MinExperience int                   `form:"min_experience" validate:"required,number"`
	LastEducation string                `form:"last_education" validate:"required"`
	Description   string                `form:"description" validate:"required"`
	CompanyEmail  string                `form:"company_email" validate:"required,email"`
	Image         *multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateRequest struct {
	ID            string                `param:"id" validate:"required,uuid"`
	JobPosition   string                `form:"job_position" validate:"omitempty"`
	JobType       string                `form:"job_type" validate:"omitempty"`
	CompanyName   string                `form:"company_name" validate:"omitempty"`
	Location      string                `form:"location" validate:"omitempty"`
	Salary        float64               `form:"salary" validate:"omitempty"`
	MinExperience int                   `form:"min_experience" validate:"omitempty,number"`
	LastEducation string                `form:"last_education" validate:"omitempty"`
	Description   string                `form:"description" validate:"omitempty"`
	CompanyEmail  string                `form:"company_email" validate:"omitempty,email"`
	Image         *multipart.FileHeader `form:"image" validate:"omitempty"`
}

type GetAllRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=newest oldest"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}