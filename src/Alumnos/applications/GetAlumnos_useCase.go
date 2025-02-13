package applications

import "practica/src/Alumnos/domain"

//import "apiRest/src/alumnos/domain"

type GetAlumnoUseCase struct {
	db domain.IAlumnoRepository
}

func NewGetAlumnoUseCase(db domain.IAlumnoRepository) *GetAlumnoUseCase {
	return &GetAlumnoUseCase{db: db}
}

func (cp *GetAlumnoUseCase) Execute(id int) (map[string]interface{}, error) {
	return cp.db.GetAlumno(id)
}