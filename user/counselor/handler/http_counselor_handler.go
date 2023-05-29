package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/usecase"
	"github.com/labstack/echo/v4"
)

type counselorHandler struct {
	CUscase usecase.CounselorUsecase
}

func NewCounselorHandler(CUcase usecase.CounselorUsecase) *counselorHandler {
	return &counselorHandler{CUscase: CUcase}
}

func (h *counselorHandler) GetAll(c echo.Context) error {

	page, _ :=  helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))
	
	page, offset, limit := helper.GetPaginateData(page, limit)

	counselors, err := h.CUscase.GetAll(offset, limit)
	
	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	totalPages, err := h.CUscase.GetTotalPages(limit)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success get all conselor", getStatusCode(err), echo.Map{
		"counselors": counselors,
		"current_pages": page,
		"total_pages": totalPages,
	}))
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	switch err {
		case counselor.ErrInternalServerError:
			return http.StatusInternalServerError
			
		case counselor.ErrCounselorNotFound:
			return http.StatusNotFound

		case 
			counselor.ErrCounselorConflict,
			counselor.ErrEmailConflict:
			return http.StatusConflict
		
		case 
		 	counselor.ErrProfilePictureFormat,
			counselor.ErrEmailFormat,
			counselor.ErrTarifFormat,
			counselor.ErrInvalidTopic,
			counselor.ErrIdFormat,
		 	counselor.ErrRequired:
			return http.StatusBadRequest

		default:
			return http.StatusInternalServerError
	}

}