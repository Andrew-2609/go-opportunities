package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

// @Summary Get Opening
// @Description Get a job opening
// @Tags Openings
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} schema.OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [get]
func GetOpeningHandler(ctx *gin.Context) {
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

	sendOK(ctx, opening)
}
