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

//manda los datos al json
func (cp *CreateAlumnoController) Execute(c *gin.Context) {
	var alumno struct {
		Nombre   string `json:"nombre"`
		Telefono string `json:"telefono"`
	}

	if err := c.BindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := strconv.Atoi(alumno.Telefono); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telefono must be a number"})
		return
	}
	cp.useCase.Execute(alumno.Nombre, alumno.Telefono)

	c.JSON(http.StatusOK, gin.H{"message": "Alumno creado exitosamente"})
}
