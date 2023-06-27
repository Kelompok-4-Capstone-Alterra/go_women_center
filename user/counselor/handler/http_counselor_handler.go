package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
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

	var getAllReq counselor.GetAllRequest

	c.Bind(&getAllReq)

	if err := isRequestValid(&getAllReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	var topicStr string
	
	if topic, ok := constant.TOPICS[getAllReq.Topic]; ok {
		topicStr = topic[0]
	}

	counselors, err := h.CUscase.GetAll(getAllReq.Search, topicStr, getAllReq.SortBy, getAllReq.IsAvailable)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all counselor", http.StatusOK, echo.Map{
		"counselors":    counselors,
	}))
}

func (h *counselorHandler) GetById(c echo.Context) error {

	var idReq counselor.IdRequest

	c.Bind(&idReq)

	if err := isRequestValid(&idReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	counselorRes, err := h.CUscase.GetById(idReq.ID)

	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseData(err.Error(), http.StatusNotFound, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get counselor by id", http.StatusOK, echo.Map{
		"counselor": counselorRes,
	}))
}

func (h *counselorHandler) GetAllReview(c echo.Context) error {

	getAllReviewReq := counselor.GetAllReviewRequest{}

	c.Bind(&getAllReviewReq)

	if err := isRequestValid(&getAllReviewReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	page, offset, limit := helper.GetPaginateData(getAllReviewReq.Page, getAllReviewReq.Limit, "mobile")
	
	reviews, totalPages, err := h.CUscase.GetAllReview(getAllReviewReq.CounselorID, offset, limit)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case counselor.ErrCounselorNotFound:
				status = http.StatusNotFound
			case counselor.ErrReviewAlreadyExist:
				status = http.StatusConflict
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	if page > totalPages {
		return c.JSON(http.StatusNotFound, helper.ResponseData(counselor.ErrPageNotFound.Error(), http.StatusNotFound, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all counselor review", http.StatusOK, echo.Map{
		"reviews":       reviews,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (h *counselorHandler) CreateReview(c echo.Context) error {

	var reviewReq counselor.CreateReviewRequest

	var user = c.Get("user").(*helper.JwtCustomUserClaims)

	c.Bind(&reviewReq)

	if err := isRequestValid(&reviewReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	reviewReq.UserID = user.ID

	err := h.CUscase.CreateReview(reviewReq)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case 
				counselor.ErrTransactionNotFound,
				counselor.ErrCounselorNotFound:
				status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create review", http.StatusOK, nil))
}
