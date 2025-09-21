// internal/handlers/auth.go
package handlers

import (
	"fmt"
	"net/http"

	"files-api/internal/config"
	"files-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	fmt.Println("Authentication successful. Generating JWT.")
	token, err := utils.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to encode JWT:", err)
		// Return a standardized error response
		c.JSON(http.StatusInternalServerError, config.SignResponse{
			Success: false,
			Message: "Failed to generate token",
			Error:   err.Error(),
		})
		return
	}

	// Return a standardized success response
	c.JSON(http.StatusOK, config.SignResponse{
		Success: true,
		Message: "Login successful",
		Data: config.JWTAccess{
			Token: token,
		},
	})
}
