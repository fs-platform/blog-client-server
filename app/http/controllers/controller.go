package controllers

type response struct {
	StatusCode uint32      `json:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}
type Controller struct {
}

func (c Controller) Format(data interface{}, statusCode uint32, message string) response {
	return response{
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
	}
}
