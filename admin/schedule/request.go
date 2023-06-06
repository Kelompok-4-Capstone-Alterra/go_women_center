package schedule

type CreateRequest struct {
	CounselorId string   `param:"id" validate:"required,uuid"`
	Dates       []string `json:"dates" validate:"required"`
	Times       []string `json:"times" validate:"required"`
}

type CounselorIdRequest struct {
	CounselorId string `param:"id" validate:"required,uuid"`
}