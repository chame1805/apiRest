package infraestructure

import (
	"net/http"
	"practica/src/Alumnos/applications"

	"github.com/gin-gonic/gin"
)

type ListAlumnosController struct {
	useCase *applications.ListAlumnosUseCase
}

func NewListAlumnosController(useCase *applications.ListAlumnosUseCase) *ListAlumnosController {
	return &ListAlumnosController{useCase: useCase}
}

func (lp_c *ListAlumnosController) Execute(c *gin.Context) {
	alumnos, err := lp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los alumnos"})
		return
	}

	c.JSON(http.StatusOK, alumnos)
}
