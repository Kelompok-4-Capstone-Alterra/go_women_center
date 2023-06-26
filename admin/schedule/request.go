package schedule

type CreateRequest struct {
	CounselorId string   `param:"id" validate:"required,uuid"`
	Dates       []string `json:"dates" validate:"gt=0,dive,required"`
	Times       []string `json:"times" validate:"gt=0,dive,required"`
}

type UpdateRequest struct {
	CounselorId string   `param:"id" validate:"required,uuid"`
	Dates       []string `json:"dates" validate:"gt=0,dive,required"`
	Times       []string `json:"times" validate:"gt=0,dive,required"`
}

type CounselorIdRequest struct {
	CounselorId string `param:"id" validate:"required,uuid"`
}