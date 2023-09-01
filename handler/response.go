package handler

import (
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

func sendOK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func sendCreated(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, data)
}

func sendNoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}
