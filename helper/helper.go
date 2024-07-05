package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseApi(message string, status string, code int, data interface{}) Response {
	metaRes := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}

	responseAPI := Response{
		Meta: metaRes,
		Data: data,
	}

	return responseAPI
}
