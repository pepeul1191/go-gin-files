// internal/middleware/detail_recovery.go
package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func DetailedRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// Log del error
		log.Printf("Panic recuperado: %v\nURL: %s\nMethod: %s",
			recovered, c.Request.URL, c.Request.Method)

		// Respuesta al cliente
		if gin.IsDebugging() {
			// En desarrollo: mostrar detalles del error
			c.JSON(500, gin.H{
				"error":   "Internal Server Error",
				"details": fmt.Sprintf("%v", recovered),
				"path":    c.Request.URL.Path,
			})
		} else {
			// En producción: respuesta genérica
			c.JSON(500, gin.H{
				"error": "Internal Server Error",
			})
		}

		c.Abort()
	})
}
