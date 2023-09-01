package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
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

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("Error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "An unexpected error occurred. Please try again later or contact the support team!")
		return
	}

	sendNoContent(ctx)
}
