package applications

import "practica/src/Alumnos/domain"

//import "apiRest/src/alumnos/domain"

type ListAlumnosUseCase struct {
	db domain.IAlumnoRepository
}

func NewListAlumnosUseCase(db domain.IAlumnoRepository) *ListAlumnosUseCase {
	return &ListAlumnosUseCase{db: db}
}

func (cp *ListAlumnosUseCase) Execute() ([]map[string]interface{}, error) {
	return cp.db.GetAll()
}