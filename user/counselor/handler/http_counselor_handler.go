package handler

import (
	"fmt"
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
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
	
	page, offset, limit := helper.GetPaginateData(page, limit, "mobile")

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

func (h *counselorHandler) GetById(c echo.Context) error {
	
	var id counselor.IdRequest

	c.Bind(&id)

	if err := isRequestValid(&id); err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	counselor, err := h.CUscase.GetById(id.ID)
	fmt.Println(err)
	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success get counselor by id", getStatusCode(err), echo.Map{
		"counselor": counselor,
	}))
}

func (h *counselorHandler) GetAllReview(c echo.Context) error {
	
	var id counselor.IdRequest

	c.Bind(&id)

	if err := isRequestValid(&id); err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	page, _ :=  helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit, "mobile")

	reviews, err := h.CUscase.GetAllReview(id.ID, offset, limit)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	totalPage, err := h.CUscase.GetTotalPagesReview(id.ID, limit)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success get all counselor review", getStatusCode(err), echo.Map{
		"reviews": reviews,
		"current_pages": page,
		"total_pages": totalPage,
	}))
}

func (h *counselorHandler) CreateReview(c echo.Context) error {
	
	var reviewReq counselor.CreateReviewRequest

	var user = c.Get("user").(*entity.UserDecodeJWT)

	c.Bind(&reviewReq)

	if err := isRequestValid(&reviewReq); err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}
	fmt.Println("user -> ", user)
	reviewReq.UserID = user.ID

	err := h.CUscase.CreateReview(reviewReq)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success create review", getStatusCode(err), nil))
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