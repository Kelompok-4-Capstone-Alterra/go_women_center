package career

type CreateRequest struct {
	JobPosition   string  `form:"jobposition" validate:"required"`
	CompanyName   string  `form:"companyname" validate:"required"`
	Location      string  `form:"location" validate:"required"`
	Salary        float64 `form:"salary" validate:"required"`
	MinExperience string  `form:"minexperience" validate:"required"`
	LastEducation string  `form:"lasteducation" validate:"required"`
	Description   string  `form:"description" validate:"required"`
	CompanyEmail  string  `form:"companyemail" validate:"required,email"`
	ApplyLink     string  `form:"applylink" validate:"required"`
}

type UpdateRequest struct {
	ID            string  `param:"id" validate:"required,uuid"`
	JobPosition   string  `form:"jobposition" validate:"omitempty"`
	CompanyName   string  `form:"companyname" validate:"omitempty"`
	Location      string  `form:"location" validate:"omitempty"`
	Salary        float64 `form:"salary" validate:"omitempty,number"`
	MinExperience string  `form:"minexperience" validate:"omitempty"`
	LastEducation string  `form:"lasteducation" validate:"omitempty"`
	Description   string  `form:"description" validate:"omitempty"`
	CompanyEmail  string  `form:"companyemail" validate:"omitempty,email"`
	ApplyLink     string  `form:"applylink" validate:"omitempty"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
