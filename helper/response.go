package helper

type responseMeta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type response struct {
	Meta responseMeta             `json:"meta"`
	Data map[string]interface{} `json:"data"`
}

func ResponseSuccess(message string, status int, data map[string]interface{}) response {
	return response{
		Meta: responseMeta{
			Message: message,
			Status:  status,
		},
		Data: data,
	}
}

func ResponseError(message string, status int) response {
	return response{
		Meta: responseMeta{
			Message: message,
			Status:  status,
		},
		Data: nil,
	}
}
