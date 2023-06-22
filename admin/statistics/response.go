package statistics

type GetStatisticsResponse struct {
	TotalUser        int `json:"total_user"`
	TotalCounselor   int `json:"total_counselor"`
	TotalTransaction int `json:"total_transaction"`
}
