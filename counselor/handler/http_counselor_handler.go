package handler

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type counselorHandler struct{
	CUscase domain.CounselorUsecase
}

func NewCounselorHandler(CUcase domain.CounselorUsecase) *counselorHandler {
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
		"current_page": page,
		"total_pages": totalPages,
	}))
}

func isRequestValid(c counselor.CreateRequest) error {

	validate := validator.New()
	err := validate.Struct(c)
	
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			
			if err.Tag() == "required" {
				return counselor.ErrRequired
			}

			if field == "email" {
				return counselor.ErrEmailFormat
			}
			
			if field == "tarif" {
				return counselor.ErrTarifFormat
			}
			
		}
	}

	if !constant.TOPIC[c.Topic] {
		return counselor.ErrInvalidTopic
	}
	
	return nil
}

func (h *counselorHandler) Create(c echo.Context) error {

	counselorReq := counselor.CreateRequest{}
	
	c.Bind(&counselorReq)

	if err := isRequestValid(counselorReq); err != nil {
		return c.JSON(
			getStatusCode(err), 
			helper.ResponseError(err.Error(),
			getStatusCode(err)),
		)
	}
	
	imgInput, _ := c.FormFile("profile_picture")

	if err := isImageValid(imgInput); err != nil {
		return c.JSON(
			getStatusCode(err),
			helper.ResponseError(err.Error(),
			getStatusCode(err)),
		)
	}

	// upload image to s3
	path, err := helper.UploadImageToS3(imgInput, "women-center")

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(counselor.ErrInternalServerError.Error(), getStatusCode(err)))
	}

	counselorReq.ProfilePicture = path

	err = h.CUscase.Create(counselorReq)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(getStatusCode(err), helper.ResponseSuccess("success create counselor", getStatusCode(err), nil))

}

func isImageValid(img *multipart.FileHeader) error {

	if img == nil {
		return counselor.ErrRequired
	}

	if !helper.IsImageValid(img) {
		return counselor.ErrProfilePictureFormat
	}

	return nil
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	switch err {
		case counselor.ErrInternalServerError:
			return http.StatusInternalServerError
			
		case counselor.ErrNotFound:
			return http.StatusNotFound

		case counselor.ErrConflict:
			return http.StatusConflict
		
		case 
		 	counselor.ErrProfilePictureFormat,
			counselor.ErrEmailFormat,
			counselor.ErrTarifFormat,
			counselor.ErrInvalidTopic,
		 	counselor.ErrRequired:
			return http.StatusBadRequest

		default:
			return http.StatusInternalServerError
	}

}