package handler

import (
	"net/http"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

// @Summary Get Openings
// @Description Get an array of job openings
// @Tags Openings
// @Produce json
// @Success 200 {array} schema.OpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error retrieving all openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "An unexpected error occurred. Please try again later or contact the support team!")
		return
	}

	sendOK(ctx, openings)
}
