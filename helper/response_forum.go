package helper

import "github.com/labstack/echo"

type responseMeta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type response struct {
	Meta responseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

func ResponseData(c echo.Context, status int, message string, data interface{}) response {
	return response{
		Meta: responseMeta{
			Message: message,
			Status:  status,
		},
		Data: data,
	}
}
