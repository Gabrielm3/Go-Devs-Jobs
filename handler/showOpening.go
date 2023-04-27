package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}