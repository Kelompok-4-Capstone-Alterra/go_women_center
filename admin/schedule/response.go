package schedule

type GetAllResponse struct {
	Dates []string `json:"dates"`
	Times []string `json:"times"`
}