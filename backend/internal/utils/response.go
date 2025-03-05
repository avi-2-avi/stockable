package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(context *gin.Context, response APIResponse) {
	context.JSON(response.Code, response)
}
