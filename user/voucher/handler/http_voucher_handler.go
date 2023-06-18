package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/usecase"
	"github.com/labstack/echo/v4"
)

type VoucherHandler interface{}

type voucherHandler struct {
	Usecase usecase.VoucherUsecase
}

func NewVoucherHandler(voucherUsecase usecase.VoucherUsecase) *voucherHandler {
	return &voucherHandler{
		Usecase: voucherUsecase,
	}
}

func (h *voucherHandler) GetAll(c echo.Context) error {
	user := c.Get("user").(*helper.JwtCustomUserClaims)
	err := user.Valid()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	data, err := h.Usecase.GetAll(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success get all voucher",
		http.StatusOK,
		data,
	))
}