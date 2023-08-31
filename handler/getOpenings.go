package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
