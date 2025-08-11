// internal/routes/routes.go
package routes

import (
	"files-api/internal/handlers"
	"files-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	// Logging
	r.Use(gin.Logger())

	// Ruta principal
	r.GET("/", handlers.Home)

	// Sign-in
	r.POST("/api/v1/sign-in", middleware.SignInAuthenticate(), handlers.SignIn)

	// Grupo de rutas protegidas
	r.POST("/api/v1/files/:folder_name", middleware.CheckJWT(), middleware.FileValidation(), handlers.UploadFile)
	r.POST("/api/v1/public/:folder_name", middleware.CheckJWT(), middleware.FileValidation(), handlers.UploadFileToPublic)
	r.GET("/api/v1/files/:folder_name/:file_name", middleware.CheckJWT(), handlers.DownloadFile)
	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Recurso no encontrado",
			"error":   c.Request.Method + " " + c.Request.URL.Path + " no existe",
		})
	})

	return r
}
