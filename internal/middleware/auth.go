// internal/middleware/auth.go
package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AuthHeader = "dev-secret-header" // Usa os.Getenv en producción

func Authenticate(c *gin.Context) {
	incoming := c.GetHeader("X-Auth-Trigger")
	if incoming != AuthHeader {
		fmt.Println("Unauthorized access attempt.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid or missing X-Auth-Trigger",
		})
		c.Abort()
		return
	}
	c.Next()
}
