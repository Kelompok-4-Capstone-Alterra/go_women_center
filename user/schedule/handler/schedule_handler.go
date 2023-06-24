package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/usecase"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	scheduleUcase usecase.ScheduleUsecase
}

func NewScheduleHandler(scheduleUcase usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{scheduleUcase}
}

func(h *ScheduleHandler) GetCurrSchedule(c echo.Context) error {

	var req schedule.GetScheduleTimeRequest

	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	scheduleTimes, err := h.scheduleUcase.GetCurrSchedule(req.CounselorId)

	if err != nil {
		status := http.StatusInternalServerError

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get current schedule", http.StatusOK, echo.Map{
		"schedule": scheduleTimes,
	}))

}