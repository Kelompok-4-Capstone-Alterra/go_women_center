package career

import "mime/multipart"

type CreateRequest struct {
	JobPosition    string                `form:"job_position" validate:"required"`
	CompanyName    string                `form:"company_name" validate:"required"`
	Location       string                `form:"location" validate:"required"`
	Salary         float64               `form:"salary" validate:"required"`
	MinExperience  string                `form:"min_experience" validate:"required"`
	LastEducation  string                `form:"last_education" validate:"required"`
	Description    string                `form:"description" validate:"required"`
	CompanyEmail   string                `form:"company_email" validate:"required,email"`
	Image *multipart.FileHeader 		 `form:"image" validate:"required"`
}

type UpdateRequest struct {
	ID            string  `param:"id" validate:"required,uuid"`
	JobPosition   string  `form:"job_position" validate:"omitempty"`
	CompanyName   string  `form:"company_name" validate:"omitempty"`
	Location      string  `form:"location" validate:"omitempty"`
	Salary        float64 `form:"salary" validate:"omitempty,number"`
	MinExperience string  `form:"min_experience" validate:"omitempty"`
	LastEducation string  `form:"last_education" validate:"omitempty"`
	Description   string  `form:"description" validate:"omitempty"`
	CompanyEmail  string  `form:"company_email" validate:"omitempty,email"`
	Image *multipart.FileHeader `form:"image" validate:"omitempty,required"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}