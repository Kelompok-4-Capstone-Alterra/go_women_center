package career

import "mime/multipart"

type CreateRequest struct {
	JobPosition   string                `form:"job_position" validate:"required"`
	CompanyName   string                `form:"company_name" validate:"required"`
	Location      string                `form:"location" validate:"required"`
	MinSalary     float64               `form:"min_salary" validate:"omitempty"`
	MaxSalary     float64               `form:"max_salary" validate:"omitempty"`
	MinExperience string                `form:"min_experience" validate:"required"`
	LastEducation string                `form:"last_education" validate:"required"`
	Description   string                `form:"description" validate:"required"`
	Requirement   string                `form:"requirement" validate:"required"`
	CompanyEmail  string                `form:"company_email" validate:"required,email"`
	Image         *multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateRequest struct {
	ID            string                `param:"id" validate:"required,uuid"`
	JobPosition   string                `form:"job_position" validate:"omitempty"`
	CompanyName   string                `form:"company_name" validate:"omitempty"`
	Location      string                `form:"location" validate:"omitempty"`
	MinSalary     float64               `form:"min_salary" validate:"omitempty,number"`
	MaxSalary     float64               `form:"max_salary" validate:"omitempty,number"`
	MinExperience string                `form:"min_experience" validate:"omitempty"`
	LastEducation string                `form:"last_education" validate:"omitempty"`
	Description   string                `form:"description" validate:"omitempty"`
	Requirement   string                `form:"requirement" validate:"omitempty"`
	CompanyEmail  string                `form:"company_email" validate:"omitempty,email"`
	Image         *multipart.FileHeader `form:"image" validate:"omitempty,required"`
}

type GetAllRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=newest highest_salary"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}
