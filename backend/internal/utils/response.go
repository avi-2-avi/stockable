package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

func Respond(context *gin.Context, response APIResponse) {
	context.JSON(response.Status, response)
}
