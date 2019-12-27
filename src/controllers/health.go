package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "alive!",
		})
	})

	return router
}
