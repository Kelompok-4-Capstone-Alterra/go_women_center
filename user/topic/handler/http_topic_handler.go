package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/topic"
	"github.com/labstack/echo/v4"
)

type topicHandler struct{}

func NewTopicHandler() *topicHandler {
	return &topicHandler{}
}

func(h *topicHandler) GetAll(c echo.Context) error{
	var topicsRes = make([]topic.GetAllResponse, 0, len(constant.TOPICS))

	for i := 1; i <= len(constant.TOPICS); i++{
		topic := topic.GetAllResponse{
			ID: i,
			Name: constant.TOPICS[i][0],
			Image: constant.TOPICS[i][1],
		}
		topicsRes = append(topicsRes, topic)
	}
	
	return c.JSON(http.StatusOK,
		helper.ResponseData("success get all topic",
		http.StatusOK,
		echo.Map{
			"topics": topicsRes,
		}),
	)
}
