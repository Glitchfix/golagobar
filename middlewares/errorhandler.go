package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorMiddleware - Errorenticate endpoints-
func ErrorMiddleware(c *gin.Context) {
	var errorMsg gin.H
	print("NOPE")
	status, _ := c.Get("status")

	switch status.(int) {
	case http.StatusBadRequest:
		errorMsg = gin.H{
			"success": false,
			"text":    "Request is nope not dope",
		}
	case http.StatusUnauthorized:
		errorMsg = gin.H{
			"success": false,
			"text":    "You shall not pass",
		}
	case http.StatusUnprocessableEntity:
		errorMsg = gin.H{
			"success": false,
			"text":    "Fimble fumble can lift the dumble",
		}
	case http.StatusInternalServerError:
		errorMsg = gin.H{
			"success": false,
			"text":    "A'ight I'mma quit",
		}
	case http.StatusNotFound:
		errorMsg = gin.H{
			"success": false,
			"text":    "Weepity woopity not on my property",
		}
	}
	c.AbortWithStatusJSON(status.(int), errorMsg)
	return
}
