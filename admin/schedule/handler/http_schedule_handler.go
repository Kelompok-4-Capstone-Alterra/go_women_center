package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule"
	Schedule "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)


type scheduleHandler struct {
	Usecase Schedule.ScheduleUsecase
}

func NewScheduleHandler(usecase Schedule.ScheduleUsecase) *scheduleHandler {
	return &scheduleHandler{
		Usecase: usecase,
	}
}

func (h *scheduleHandler) GetByCounselorId(c echo.Context) error {
	
	var req schedule.CounselorIdRequest

	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	schedule, err := h.Usecase.GetByCounselorId(req.CounselorId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get schedule counselor", http.StatusOK, echo.Map{
		"schedule": schedule,
	}))
}

func (h *scheduleHandler) Create(c echo.Context) error {

	req := schedule.CreateRequest{}

	c.Bind(&req)	

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}
	

	err := h.Usecase.Create(req)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case schedule.ErrCounselorNotFound:
				status = http.StatusNotFound
			case schedule.ErrTimeInvalid,
				schedule.ErrDateInvalid:
				status = http.StatusBadRequest
		}

		return c.JSON(status,
			helper.ResponseData(err.Error(), status, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create schedule counselor", http.StatusOK, nil))
}

func(h *scheduleHandler) Update(c echo.Context) error {

	req := schedule.UpdateRequest{}

	c.Bind(&req)
	
	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.Usecase.Update(req)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case schedule.ErrCounselorNotFound:
				status = http.StatusNotFound
			case schedule.ErrTimeInvalid,
				schedule.ErrDateInvalid:
				status = http.StatusBadRequest
		}

		return c.JSON(status,
			helper.ResponseData(err.Error(), status, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update schedule counselor", http.StatusOK, nil))
}

func(h *scheduleHandler) Delete(c echo.Context) error {

	var req schedule.CounselorIdRequest

	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.Usecase.Delete(req.CounselorId)

	if err != nil {
		status := http.StatusInternalServerError
		if err == schedule.ErrCounselorNotFound{
			status = http.StatusNotFound
		}
		return c.JSON(status,
			helper.ResponseData(err.Error(), status, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete schedule counselor", http.StatusOK, nil))
}