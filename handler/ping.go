package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Ping - ping check for server
func Ping(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"succes": true,
		"time":   time.Now().String(),
	})
}
