package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrew-2609/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

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
