package infraestructure

import (
	"net/http"
	"practica/src/Materias/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMateriaController struct {
	useCase *applications.DeleteMateriaUseCase
}

func NewDeleteMateriaController(useCase *applications.DeleteMateriaUseCase) *DeleteMateriaController {
	return &DeleteMateriaController{useCase: useCase}
}

func (dr_c *DeleteMateriaController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dr_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la materia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada exitosamente"})
}