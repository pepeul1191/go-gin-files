// internal/config/config.go
package config

import "os"

var (
	JWTSecret  = os.Getenv("JWT_SECRET")
	AuthHeader = os.Getenv("AUTH_HEADER")
)

func init() {
	if JWTSecret == "" {
		JWTSecret = "default_jwt_secret"
	}
	if AuthHeader == "" {
		AuthHeader = "dev-secret-header"
	}
}
