package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseFormat struct {
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

func SendSuccess(ctx *gin.Context, Status int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"ok":      false,
		"message": msg,
		"status":  code,
	})
}
