package applications

import "practica/src/Alumnos/domain"

//import "apiRest/src/alumnos/domain"

type DeleteAlumnoUseCase struct {
 db domain.IAlumnoRepository
}

func NewDeleteAlumnoUseCase(db domain.IAlumnoRepository) *DeleteAlumnoUseCase {
	return &DeleteAlumnoUseCase{db: db}
}

func (cp *DeleteAlumnoUseCase) Execute(id int) error {
	return cp.db.Delete(id)
}