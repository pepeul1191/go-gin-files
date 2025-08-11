// internal/middleware/auth.go
package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"files-api/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignInAuthenticate(c *gin.Context) {
	incoming := c.GetHeader("X-Auth-Trigger")
	if incoming != config.AuthHeader {
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

// CheckJWT verifica y valida un token JWT en el encabezado Authorization
func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extraer el token del encabezado Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			return
		}

		// 2. Verificar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid Authorization header format (expected: Bearer <token>)",
			})
			return
		}

		tokenString := parts[1]

		// 3. Parsear y validar el token usando JWTSecret de config
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar el algoritmo de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWTSecret), nil
		}, jwt.WithLeeway(5*time.Second))

		// 4. Manejar errores de validación
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "Malformed token",
				})
			case errors.Is(err, jwt.ErrTokenExpired):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "Token has expired",
				})
			case errors.Is(err, jwt.ErrTokenNotValidYet):
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "Token not active yet",
				})
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "Token validation failed",
				})
			}
			return
		}

		// 5. Verificar claims del token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Almacenar claims en el contexto para uso posterior
			c.Set("jwtClaims", claims)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token claims",
			})
			return
		}

		c.Next()
	}
}
