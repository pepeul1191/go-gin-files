// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"files-api/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Configurar rutas
	r := routes.Setup()

	// Servir archivos estáticos
	r.Static("/uploads", "uploads")

	// Iniciar servidor
	fmt.Printf("Servidor escuchando en http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
