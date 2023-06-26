package schedule

type GetScheduleTimeRequest struct {
	CounselorId string `param:"id" validate:"required,uuid"`
}