package infraestructure

import (
	"net/http"
	"practica/src/Alumnos/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteAlumnoController struct {
	useCase *applications.DeleteAlumnoUseCase
}

func NewDeleteAlumnoController(useCase *applications.DeleteAlumnoUseCase) *DeleteAlumnoController {
	return &DeleteAlumnoController{useCase: useCase}
}

func (dr_c *DeleteAlumnoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dr_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el alumno"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumno eliminado exitosamente"})
}
