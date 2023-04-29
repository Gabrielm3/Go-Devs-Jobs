package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gabrielm3/Go-Devs-Jobs/schemas"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Sucess 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error llisting openings")
		return
	}

	sendSucess(ctx, "list-opening", openings)
}
