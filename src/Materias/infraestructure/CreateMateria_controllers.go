package infraestructure

import (
	"net/http"
	"practica/src/Materias/applications"

	"github.com/gin-gonic/gin"
)
//hace referencia al use case para poder crear como tal la materia
type CreateMateriaController struct {
	useCase *applications.CreateMateriaUseCase
}
//un contructor que recibe el use case
func NewCreateMateriaController(useCase *applications.CreateMateriaUseCase) *CreateMateriaController {
	return &CreateMateriaController{useCase: useCase}
}

//manda los datos al json
func (cp *CreateMateriaController) Execute(c *gin.Context) {
	var materia struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&materia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cp.useCase.Execute(materia.Name)

	c.JSON(http.StatusOK, gin.H{"message": "Materia creada exitosamente"})
}