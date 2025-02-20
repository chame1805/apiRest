package infraestructure

import (
	"net/http"
	"practica/src/Materias/applications"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMateriaByIDController struct {
	useCase *applications.GetMateriaUseCase
}

func NewGetMateriaController(useCase *applications.GetMateriaUseCase) *GetMateriaByIDController {
	return &GetMateriaByIDController{useCase: useCase}
}

func (gp_c *GetMateriaByIDController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	materia, err := gp_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la materia"})
		return
	}

	if materia == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
		return
	}

	c.JSON(http.StatusOK, materia)
}