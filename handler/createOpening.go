package handler

import (
	"net/http"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Request validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role:      request.Role,
		Company:   request.Company,
		Location:  request.Location,
		WorkModel: request.WorkModel,
		Link:      request.Link,
		Salary:    request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "An unexpected error occurred. Please try again later or contact the support team!")
		return
	}

	sendCreated(ctx, opening)
}
