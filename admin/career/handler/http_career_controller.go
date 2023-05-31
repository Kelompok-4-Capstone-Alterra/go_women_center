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
		career.ErrImageFormat,
		career.ErrEmailFormat,
		career.ErrImageSize,
		career.ErrIdFormat,
		career.ErrRequired:
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

func (h *careerHandler) Create(c echo.Context) error {

	careerReq := career.CreateRequest{}

	c.Bind(&careerReq)

	if err := isRequestValid(careerReq); err != nil {

		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	imgInput, _ := c.FormFile("image")

	if err := isImageValid(imgInput); err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	err := h.CareerUsecase.Create(careerReq, imgInput)

	if err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success create career", getStatusCode(err), nil))

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

func (h *careerHandler) Update(c echo.Context) error {

	var careerReq career.UpdateRequest

	c.Bind(&careerReq)
	 
	if err := isRequestValid(careerReq); err != nil {	
		return c.JSON(
			getStatusCode(err), 
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	imgInput, _ := c.FormFile("image")

	if err := isImageValid(imgInput); err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	err := h.CareerUsecase.Update(careerReq, imgInput)

	if err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success update career", getStatusCode(err), nil))
}

func (h *careerHandler) Delete(c echo.Context) error {

	var id career.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	err := h.CareerUsecase.Delete(id.ID)

	if err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(), getStatusCode(err)),
		)
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success delete career", getStatusCode(err), nil))

}
