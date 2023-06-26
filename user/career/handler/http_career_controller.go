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

	var getAllReq career.GetAllRequest

	c.Bind(&getAllReq)

	if err := isRequestValid(&getAllReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}	

	careers, err := h.CareerUsecase.GetAll(getAllReq.Search, getAllReq.SortBy)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all career", http.StatusOK, echo.Map{
		"careers":       careers,
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

	careerRes, err := h.CareerUsecase.GetById(id.ID)

	if err != nil {

		status := http.StatusInternalServerError

		if err == career.ErrCareerNotFound {
			status = http.StatusNotFound
		}

		return c.JSON(
			status,
			helper.ResponseData(err.Error(), status, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get career by id", http.StatusOK, careerRes))

}
