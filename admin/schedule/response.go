package schedule

type GetAllResponse struct {
	Date []string `json:"dates"`
	Time []string `json:"times"`
}