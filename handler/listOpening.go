package handler

import (
	"net/http"

	"github.com/Gabrielm3/Go-Devs-Jobs/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error llisting openings")
		return
	}

	sendSucess(ctx, "list-opening", openings)
}
