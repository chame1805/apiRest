package infraestructure

import (
	"net/http"
	"practica/src/Alumnos/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAlumnoByIDController struct {
	useCase *applications.GetAlumnoUseCase
}

func NewGetAlumnoController(useCase *applications.GetAlumnoUseCase) *GetAlumnoByIDController {
	return &GetAlumnoByIDController{useCase: useCase}
}

func (gp_c *GetAlumnoByIDController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	alumno, err := gp_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el alumno"})
		return
	}

	if alumno == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
		return
	}

	c.JSON(http.StatusOK, alumno)
}
