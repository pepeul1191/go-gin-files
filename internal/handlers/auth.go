// internal/handlers/auth.go
package handlers

import (
	"files-api/internal/middleware"
	"fmt"
	"net/http"

	"files-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	incoming := c.GetHeader("X-Auth-Trigger")
	if incoming != middleware.AuthHeader {
		fmt.Println("Unauthorized access attempt.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid or missing X-Auth-Trigger",
		})
		return
	}

	fmt.Println("Authentication successful. Generating JWT.")
	token, err := utils.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to encode JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
