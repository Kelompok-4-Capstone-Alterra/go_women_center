package handler

import (
	"net/http"
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type counselorHandler struct{
	CUscase usecase.CounselorUsecase
}

func NewCounselorHandler(CUcase usecase.CounselorUsecase) *counselorHandler {
	return &counselorHandler{CUscase: CUcase}
}

func(h *counselorHandler) GetAll(c echo.Context) error {

	page, _ :=  strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	search := c.QueryParam("search")
	
	page, offset, limit := helper.GetPaginateData(page, limit)
	
	counselors, totalPages, err := h.CUscase.GetAll(search, offset, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	if page > totalPages {
		return c.JSON(http.StatusNotFound, helper.ResponseData(counselor.ErrPageNotFound.Error(), http.StatusNotFound, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all conselor", http.StatusOK, echo.Map{
		"counselors": counselors,
		"current_pages": page,
		"total_pages": totalPages,
	}))
}

func(h *counselorHandler) Create(c echo.Context) error {

	counselorReq := counselor.CreateRequest{}
	file, _ :=  c.FormFile("profile_picture")
	counselorReq.ProfilePicture = file

	c.Bind(&counselorReq)

	if err := isRequestValid(counselorReq); err != nil {
		return c.JSON(
			http.StatusBadRequest, 
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CUscase.Create(counselorReq)

	if err != nil {
		status := http.StatusInternalServerError

		switch err.Error() {
			case counselor.ErrEmailConflict.Error(),
				counselor.ErrUsernameConflict.Error():
				status = http.StatusConflict
			case counselor.ErrProfilePictureFormat.Error():
				status = http.StatusBadRequest
		}
		
		return c.JSON(
			status,
			helper.ResponseData(err.Error(), status, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create counselor", http.StatusOK, nil))

}

func(h *counselorHandler) GetById(c echo.Context) error {

	var id counselor.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	counselor, err := h.CUscase.GetById(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil),
	)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get counselor by id", http.StatusOK, echo.Map{
		"counselor": counselor,
	}))

}

func(h *counselorHandler) Update(c echo.Context) error {

	counselorReq := counselor.UpdateRequest{}
	file, _ :=  c.FormFile("profile_picture")
	counselorReq.ProfilePicture = file

	c.Bind(&counselorReq)
	 
	if err := isRequestValid(counselorReq); err != nil {	
		return c.JSON(
			http.StatusBadRequest, 
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CUscase.Update(counselorReq)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update counselor", http.StatusOK, nil))
}

func(h *counselorHandler) Delete(c echo.Context) error {

	var id counselor.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.CUscase.Delete(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete counselor", http.StatusOK, nil))

}
