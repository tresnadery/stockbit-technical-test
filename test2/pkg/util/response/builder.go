package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Success bool   `json:"success" default:"true"`
	Message string `json:"message" default:"true"`
}

type response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Response struct {
	Response response `json:"response"`
	Code     int      `json:"code"`
}

func ResponseBuilder(res *Response, data interface{}) *Response {
	res.Response.Data = data
	return res
}

func CustomBuilder(code int, data interface{}, message string) *Response {
	success := true
	if code != http.StatusOK {
		success = false
	}
	return &Response{
		Response: response{
			Meta: Meta{
				Success: success,
				Message: message,
			},
			Data: data,
		},
		Code: code,
	}
}

func (s *Response) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}
