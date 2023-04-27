package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
