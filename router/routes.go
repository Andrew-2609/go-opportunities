package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Openings",
			})
		})
		v1.GET("/openings/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})
		v1.PUT("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})
		v1.DELETE("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})
	}
}
