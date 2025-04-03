package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	alumnos "practica/src/Alumnos/infraestructure"
	materias "practica/src/Materias/infraestructure"
	"time"

	"github.com/gin-contrib/cors" // Importar middleware de CORS
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes, puedes cambiarlo por tu frontend: "http://localhost:3000"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rutas
	alumnos.SetupAlumnos(r)
	materias.SetupMaterias(r)

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
