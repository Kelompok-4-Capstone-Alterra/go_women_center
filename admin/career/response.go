package career

type GetAllResponse struct {
	ID           string  `json:"id"`
	Image        string  `json:"image"`
	JobPosition  string  `json:"job_position"`
	CompanyName  string  `json:"company_name"`
	Location     string  `json:"location"`
	MinSalary    float64 `json:"min_salary"`
	MaxSalary    float64 `json:"max_salary"`
	CompanyEmail string  `json:"company_email"`
}

type GetByResponse struct {
	ID            string  `json:"id"`
	Image         string  `json:"image"`
	JobPosition   string  `json:"job_position"`
	CompanyName   string  `json:"company_name"`
	Location      string  `json:"location"`
	MinSalary     float64 `json:"min_salary"`
	MaxSalary     float64 `json:"max_salary"`
	MinExperience string  `json:"min_experience"`
	LastEducation string  `json:"last_education"`
	Description   string  `json:"description"`
	Requirement   string  `json:"requirement"`
	CompanyEmail  string  `json:"company_email" validate:"required,email"`
}
