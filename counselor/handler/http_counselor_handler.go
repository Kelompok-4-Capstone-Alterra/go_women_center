package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)


type ResponseError struct {
	Message string `json:"message"`
}

type counselorHandler struct{
	CUscase domain.CounselorUsecase
}

func NewCounselorHandler() *counselorHandler {
	return &counselorHandler{}
}

func (h *counselorHandler) GetAll(c echo.Context) error {

	page, _ :=  helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))
	
	counselors, err := h.CUscase.GetAll(page, limit)
	
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(getStatusCode(err), counselors)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}

}