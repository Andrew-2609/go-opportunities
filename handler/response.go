package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"error": struct {
			Msg  string `json:"message"`
			Code int    `json:"code"`
		}{Msg: msg, Code: code},
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data any) {
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("Operation %s was successfull!", op),
		"data":    data,
	})
}
