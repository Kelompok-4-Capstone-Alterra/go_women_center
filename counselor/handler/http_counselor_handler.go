package handler

import (
	"errors"
	"net/http"
	"strings"

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
	
	counselors, err := h.CUscase.GetAll(page, limit)
	
	if err != nil {
		status := getStatusCode(err)
		return c.JSON(status, helper.ResponseError(err.Error(), status))
	}

	return c.JSON(getStatusCode(err), counselors)
}

func isRequestValid(m interface{}) error {
	validate := validator.New()
	err := validate.Struct(m)
	
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			
			if err.Tag() == "required" {
				return errors.New(field + " is required")
			}

			if err.Tag() == "email" {
				return errors.New(field + " must be a valid email")
			}
			
			if err.Tag() == "number" {
				return errors.New(field + " must be a number")
			}
			
		}
	}
	return nil
}

func (h *counselorHandler) Create(c echo.Context) error {

	counselorReq := counselor.CreateRequest{}
	
	c.Bind(&counselorReq)

	if err := isRequestValid(counselorReq); err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}
	
	imgInput, err := c.FormFile("profile_picture")

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to profile picture is required")
	}

	if !helper.IsImageValid(imgInput) {
		return c.JSON(http.StatusBadRequest, "profile picture must be an image and png/jpg/jpeg format")
	}

	// upload image to s3
	path, err := helper.UploadImageToS3(imgInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to upload image")
	}

	counselorReq.ProfilePicture = path

	err = h.CUscase.Create(counselorReq)

	if err != nil {
		return c.JSON(getStatusCode(err), helper.ResponseError(err.Error(), getStatusCode(err)))
	}

	return c.JSON(http.StatusOK, "success create counselor")

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
	default:
		return http.StatusInternalServerError
	}

}