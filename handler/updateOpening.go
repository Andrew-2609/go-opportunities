package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Request validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")

	validId, err := strconv.Atoi(id)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("ID attribute must be of type int64. '%s' is not valid!", id))
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, validId).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("There's no Opening with the id %d", validId))
		return
	}

	bindFields(request, &opening)

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "An unexpected error occurred. Please try again later or contact the support team!")
		return
	}

	sendOK(ctx, opening)
}

func bindFields(request UpdateOpeningRequest, opening *schema.Opening) {
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.WorkModel != "" {
		opening.WorkModel = request.WorkModel
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
}
