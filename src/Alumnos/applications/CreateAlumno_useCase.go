package applications

import "practica/src/Alumnos/domain"

//import "apiRest/src/Alumnos/domain"

type CreateAlumnoUseCase struct {
	db domain.IAlumnoRepository
}

func NewCreateAlumnoUseCase(db domain.IAlumnoRepository) *CreateAlumnoUseCase {
	return &CreateAlumnoUseCase{db: db}
}
// ase referencia al metodo save y pasa los parametros nombre y telefono
func (cp *CreateAlumnoUseCase) Execute(nombre string, telefono string) {
	cp.db.Save(nombre, telefono)
}
