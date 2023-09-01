package handler

import (
	"fmt"
	"net/http"

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

func sendOK(ctx *gin.Context, op string, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation %s was successfull!", op),
		"data":    data,
	})
}

func sendCreated(ctx *gin.Context, op string, data any) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Operation %s was successfull!", op),
		"data":    data,
	})
}

func sendNoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}
