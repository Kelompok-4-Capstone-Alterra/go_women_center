package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type careerHandler struct {
	CareerUsecase usecase.CareerUsecase
}

func NewCareerHandler(CareerUsecase usecase.CareerUsecase) *careerHandler {
	return &careerHandler{CareerUsecase: CareerUsecase}
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	switch err {
	case career.ErrInternalServerError:
		return http.StatusInternalServerError

	case career.ErrCareerNotFound:
		return http.StatusNotFound

	case
		career.ErrIdFormat:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

func (h *careerHandler) GetAll(c echo.Context) error {

	page, _ := helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit)

	careers, err := h.CareerUsecase.GetAll(offset, limit)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	totalPages, err := h.CareerUsecase.GetTotalPages(limit)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success get all conselor", getStatusCode(err), echo.Map{
		"careers":       careers,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (h *careerHandler) GetById(c echo.Context) error {

	var id career.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	career, err := h.CareerUsecase.GetById(id.ID)

	if err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success get career by id", getStatusCode(err), career))

}
