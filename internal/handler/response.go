package handler

import "log"

const (
	success = "success"
	err     = "error"
)

type (
	Response struct {
		Status        string           `json:"status"`
		Message       string           `json:"message"`
		ErrorResponse []*ErrorResponse `json:"errors"`
		Data          interface{}      `json:"data"`
	}

	ErrorResponse struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func GetResponse(msg string, data interface{}) *Response {
	resp := &Response{
		Status:  success,
		Message: msg,
		Data:    data,
	}
	return resp
}

func GetErrorResponse(msg string, errs []error) *Response {
	log.Println(msg)
	resp := &Response{
		Status:  err,
		Message: msg,
		Data:    nil,
	}
	if len(errs) == 0 {
		return resp
	}

	for _, err := range errs {
		if err == nil {
			continue
		}
		log.Println(err.Error())
		resp.ErrorResponse = append(resp.ErrorResponse, &ErrorResponse{
			Message: err.Error(),
		})
	}

	return resp
}
