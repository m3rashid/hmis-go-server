package controllers

import (
	"net/http"
)


func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}
