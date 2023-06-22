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

func (h *careerHandler) GetAll(c echo.Context) error {

	var getAllReq career.GetAllRequest

	c.Bind(&getAllReq)

	if err := isRequestValid(&getAllReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	page := getAllReq.Page
	limit := getAllReq.Limit

	page, offset, limit := helper.GetPaginateData(page, limit)
	careers, totalPages, err := h.CareerUsecase.GetAll(getAllReq.Search, getAllReq.SortBy, offset, limit)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	if page > totalPages {
		return c.JSON(
			http.StatusNotFound,
			helper.ResponseData(career.ErrPageNotFound.Error(), http.StatusBadRequest, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all career", http.StatusOK, echo.Map{
		"careers":       careers,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (h *careerHandler) Create(c echo.Context) error {

	careerReq := career.CreateRequest{}
	imgInput, _ := c.FormFile("image")
	careerReq.Image = imgInput
	c.Bind(&careerReq)
	helper.RemoveWhiteSpace(careerReq)

	if err := isRequestValid(careerReq); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CareerUsecase.Create(careerReq, imgInput)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create career", http.StatusOK, nil))

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

	return c.JSON(http.StatusOK, helper.ResponseData("success get career by id", http.StatusOK, career))
}

func (h *careerHandler) Update(c echo.Context) error {

	var careerReq career.UpdateRequest
	imgInput, _ := c.FormFile("image")
	careerReq.Image = imgInput
	c.Bind(&careerReq)
	helper.RemoveWhiteSpace(careerReq)

	if err := isRequestValid(careerReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CareerUsecase.Update(careerReq, imgInput)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update career", http.StatusOK, nil))
}

func (h *careerHandler) Delete(c echo.Context) error {

	var id career.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CareerUsecase.Delete(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete career", http.StatusOK, nil))

}
