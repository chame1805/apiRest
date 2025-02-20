package infraestructure

import (
	"practica/src/Materias/applications"

	"github.com/gin-gonic/gin"
)

func SetupMaterias(r *gin.Engine) {
	ps := NewMySQLMaterias()

	createMateriaController := NewCreateMateriaController(applications.NewCreateMateriaUseCase(ps))
	listMateriasController := NewListaMateriasController(applications.NewListMateriasUseCase(ps))
	getMateriaByIDController := NewGetMateriaController(applications.NewGetMateriaUseCase(ps))
	updateMateriaController := NewUpdateMateriasController(applications.NewUpdateMateriaUseCase(ps))
	deleteMateriaController := NewDeleteMateriaController(applications.NewDeleteMateriaUseCase(ps))

	r.POST("/crearmaterias", createMateriaController.Execute)
	r.GET("/materias", listMateriasController.Execute)
	r.GET("/materias/:id", getMateriaByIDController.Execute)
	r.PUT("/materias/:id", updateMateriaController.Execute)
	r.DELETE("/materias/:id", deleteMateriaController.Execute)
}