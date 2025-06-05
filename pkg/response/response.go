package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessReponse(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message[code],
		Data:    data,
	})
}

func ErrorReponse(c *gin.Context, code int, mess string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message[code],
		Data:    nil,
	})
}
