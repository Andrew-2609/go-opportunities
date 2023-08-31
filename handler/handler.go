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

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "POST Opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
