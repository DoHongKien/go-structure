package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   // status code
	Err    string      `json:"error"`  // thong bao loi
	Detail interface{} `json:"detail"` // du lai return
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	if message == "" {
		message = msg[code]
	}
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    message,
	})
}
