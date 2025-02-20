package infraestructure

import (
	"net/http"
	"practica/src/Materias/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMateriasController struct {
	useCase *applications.UpdateMateriaUseCase
}

func NewUpdateMateriasController(useCase *applications.UpdateMateriaUseCase) *UpdateMateriasController {
	return &UpdateMateriasController{useCase: useCase}
}

func (up_c *UpdateMateriasController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = up_c.useCase.Execute(id, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la materia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Materia actualizada exitosamente"})
}