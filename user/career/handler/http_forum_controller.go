package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type careerHandler struct {
	CareerUsecase usecase.CareerUsecase
}

func NewCareerHandler(CareerUsecase usecase.CareerUsecase) *careerHandler {
	return &careerHandler{CareerUsecase: CareerUsecase}
}

func (h *careerHandler) GetAll(c echo.Context) error {

	page, _ := helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit)

	careers, err := h.CareerUsecase.GetAll(offset, limit)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	totalPages, err := h.CareerUsecase.GetTotalPages(limit)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	return c.JSON(http.StatusAccepted, helper.ResponseData("success get all conselor", http.StatusAccepted, echo.Map{
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
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	career, err := h.CareerUsecase.GetById(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusAccepted, helper.ResponseData("success get career by id", http.StatusAccepted, career))

}

func (h *careerHandler) GetBySearch(c echo.Context) error {

	var search career.SearchRequest

	c.Bind(&search)

	careers, err := h.CareerUsecase.GetBySearch(search.Search)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusAccepted, helper.ResponseData("success get career by search", http.StatusAccepted, careers))
}
