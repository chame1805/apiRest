package applications

import "practica/src/Alumnos/domain"

//import "apiRest/src/alumnos/domain"

type UpdateAlumnoUseCase struct {
	db domain.IAlumnoRepository
}

func NewUpdateAlumnoUseCase(db domain.IAlumnoRepository) *UpdateAlumnoUseCase {
	return &UpdateAlumnoUseCase{db: db}
}

func (cp *UpdateAlumnoUseCase) Execute(id int, nombre string, telefono string) error {
	return cp.db.Update(id, nombre, telefono)
}
