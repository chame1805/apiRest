// package infraestructure
package infraestructure

import (
	"net/http"
	"practica/src/Alumnos/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateAlumnoController struct {
	useCase *applications.UpdateAlumnoUseCase
}

func NewUpdateAlumnoController(useCase *applications.UpdateAlumnoUseCase) *UpdateAlumnoController {
	return &UpdateAlumnoController{useCase: useCase}
}

func (up_c *UpdateAlumnoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Nombre   string `json:"nombre"`
		Telefono string `json:"telefono"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = up_c.useCase.Execute(id, input.Nombre, input.Telefono)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el alumno"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumno actualizado exitosamente"})
}
