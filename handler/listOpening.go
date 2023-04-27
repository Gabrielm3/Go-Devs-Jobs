package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
