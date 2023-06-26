package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type StatisticController interface {}

type statisticController struct {
	StatisticUsecase usecase.StatisticUsecase
}

func NewStatisticController(uc usecase.StatisticUsecase) *statisticController {
	return &statisticController{
		StatisticUsecase: uc,
	}
}

func (sc *statisticController) GetData(c echo.Context) error {
	res, err := sc.StatisticUsecase.GetStatistic()
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK,
		helper.ResponseData("success getting statistics data",
		http.StatusOK, 
		res,
	))
}