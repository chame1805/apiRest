package infraestructure

import (
	"net/http"
	"practica/src/Materias/applications"
	

	//"practica/src/Materias/applications"

	"github.com/gin-gonic/gin"
)

type ListaMateriasController struct {
	useCase *applications.ListMateriasUseCase
}


func NewListaMateriasController(useCase *applications.ListMateriasUseCase) *ListaMateriasController {
	return &ListaMateriasController{useCase: useCase}
}

func (lp_c *ListaMateriasController) Execute(c *gin.Context) {
	materias, err := lp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las materias"})
		return
	}

	c.JSON(http.StatusOK, materias)
}