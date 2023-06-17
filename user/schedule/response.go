package schedule

type Time struct {
	ID          string `json:"id"`
	Time        string `json:"time"`
	IsAvailable bool   `json:"is_available"`
}

type GetScheduleResponse struct {
	DateId string `json:"date_id"`
	Times []Time `json:"times"`
}
