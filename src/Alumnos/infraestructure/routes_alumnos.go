package infraestructure

import (

	"practica/src/Alumnos/applications"

	"github.com/gin-gonic/gin"
)

func SetupAlumnos(r *gin.Engine) {
	ps := NewMySQLAlumnos()

	createAlumnoController := NewCreateAlumnoController(applications.NewCreateAlumnoUseCase(ps))
	listAlumnosController := NewListAlumnosController(applications.NewListAlumnosUseCase(ps))
	getAlumnoByIDController := NewGetAlumnoController(applications.NewGetAlumnoUseCase(ps))
	updateAlumnoController := NewUpdateAlumnoController(applications.NewUpdateAlumnoUseCase(ps))
	deleteAlumnoController := NewDeleteAlumnoController(applications.NewDeleteAlumnoUseCase(ps))

	r.POST("/crearalumnos", createAlumnoController.Execute)
	r.GET("/alumnos", listAlumnosController.Execute)
	r.GET("/alumnos/:id", getAlumnoByIDController.Execute)
	r.PUT("/alumnos/:id", updateAlumnoController.Execute)
	r.DELETE("/alumnos/:id", deleteAlumnoController.Execute)
}
