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
	r.POST("/api/v1/sign-in", handlers.SignIn)

	// Grupo de rutas protegidas
	fileGroup := r.Group("/api/v1/files")
	fileGroup.Use(middleware.Authenticate)
	{
		fileGroup.POST("/:issue_id", handlers.UploadFile)
		fileGroup.GET("/:issue_id/:file_name", handlers.DownloadFile)
	}

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Recurso no encontrado",
			"error":   c.Request.Method + " " + c.Request.URL.Path + " no existe",
		})
	})

	return r
}
