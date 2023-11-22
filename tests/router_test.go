package tests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  200,
			"message": "pong",
		})
	})
	router.Run(":8080")
}
