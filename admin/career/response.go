package career

type GetAllResponse struct {
	ID           string  `json:"id"`
	Image        string  `json:"image"`
	JobPosition  string  `json:"job_position"`
	CompanyName  string  `json:"company_name"`
	Location     string  `json:"location"`
	Salary       float64 `json:"salary"`
	CompanyEmail string  `json:"company_email"`
}

type GetByResponse struct {
	ID            string  `json:"id"`
	Image         string  `json:"image"`
	JobPosition   string  `json:"job_position"`
	CompanyName   string  `json:"company_name"`
	Location      string  `json:"location"`
	Salary        float64 `json:"salary" validate:"omitempty,number"`
	MinExperience string  `json:"min_experience"`
	LastEducation string  `json:"last_education"`
	Description   string  `json:"description"`
	CompanyEmail  string  `json:"company_email" validate:"required,email"`
}
