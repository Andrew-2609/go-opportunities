package router

import (
	docs "github.com/Andrew-2609/go-opportunities/docs"
	"github.com/Andrew-2609/go-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.GetOpeningsHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	}

	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
