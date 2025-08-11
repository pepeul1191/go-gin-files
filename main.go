// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"files-api/internal/config"
	"files-api/internal/routes"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Advertencia: No se encontró archivo .env o no se pudo cargar")
		// Puede continuar, usará valores por defecto
	}

	// Ahora puedes acceder a las variables de entorno
	port := os.Getenv("PORT")
	fmt.Println("1 +++++++++++++++++++++++++++++")
	fmt.Println("Puerto:", port)
	fmt.Println("2 +++++++++++++++++++++++++++++")

	if port == "" {
		port = "3000" // valor por defecto
	}

	// Asegurarse de que config cargue las variables (si usas config.Init())
	config.Init() // Opcional: si quieres inicializar config en un paquete

	// Configurar rutas
	r := routes.Setup()

	// Servir archivos estáticos
	r.Static("/uploads", "uploads")

	// Iniciar servidor
	fmt.Printf("🚀 Servidor escuchando en http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}
