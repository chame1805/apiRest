package infraestructure

import (
	"net/http"
	"practica/src/Alumnos/applications"

	//"apiRest/src/alumnos/applications" // Ruta corregida

	"strconv"

	"github.com/gin-gonic/gin"
)
//hace referencia al use case para poder crear como tal el alumno
type CreateAlumnoController struct {
	useCase *applications.CreateAlumnoUseCase
}

//un contructor que recibe el use case
func NewCreateAlumnoController(useCase *applications.CreateAlumnoUseCase) *CreateAlumnoController {
	return &CreateAlumnoController{useCase: useCase}
}
func (cp *CreateAlumnoController) Execute(c *gin.Context) {
	var alumno struct {
		Nombre   string `json:"nombre"`
		Telefono string `json:"telefono"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := strconv.Atoi(alumno.Telefono); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telefono must be a number"})
		return
	}

	// Ejecutar el caso de uso con la contraseña encriptada
	err := cp.useCase.Execute(alumno.Nombre, alumno.Telefono, alumno.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumno creado exitosamente"})
}
