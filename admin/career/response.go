package career

type GetAllResponse struct {
	ID            string  `json:"id"`
	Image         string  `json:"image"`
	JobPosition   string  `form:"jobposition"`
	CompanyName   string  `form:"companyname"`
	Location      string  `form:"location"`
	Salary        float64 `form:"salary"`
	CompanyEmail  string  `form:"companyemail"`
	ApplyLink     string  `form:"applylink"`
}

type GetByResponse struct {
	ID            string  `json:"id"`
	Image         string  `json:"image"`
	JobPosition   string  `form:"jobposition"`
	CompanyName   string  `form:"companyname"`
	Location      string  `form:"location"`
	Salary        float64 `form:"salary" validate:"omitempty,number"`
	MinExperience string  `form:"minexperience"`
	LastEducation string  `form:"lasteducation"`
	Description   string  `form:"description"`
	CompanyEmail  string  `form:"companyemail" validate:"required,email"`
	ApplyLink     string  `form:"applylink" validate:"required"`
}
