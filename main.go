
package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	alumnos "practica/src/Alumnos/infraestructure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	r := gin.Default()
	alumnos.SetupAlumnos(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
