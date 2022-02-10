package response

import "github.com/gin-gonic/gin"

func JSON(ginContext *gin.Context, httpStatusCode int, statusCode int, Message string, data interface{}) {
	type Payload struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	payload := Payload{
		Code:    statusCode,
		Message: Message,
		Data:    data,
	}
	ginContext.PureJSON(httpStatusCode, payload)
}
