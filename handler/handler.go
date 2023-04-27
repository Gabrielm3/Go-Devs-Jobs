package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
func ShowOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
func DeleteOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
func UpdateOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
func ListOpeningHandler() {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
