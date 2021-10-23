package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "Health check server OK"})
}
