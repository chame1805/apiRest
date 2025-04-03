package applications

import "practica/src/Alumnos/domain"

type CreateAlumnoUseCase struct {
	db domain.IAlumnoRepository
}

func NewCreateAlumnoUseCase(db domain.IAlumnoRepository) *CreateAlumnoUseCase {
	return &CreateAlumnoUseCase{db: db}
}

func (cp *CreateAlumnoUseCase) Execute(nombre string, telefono string, password string) error {
	// Hashear la contraseña antes de guardarla
	hashedPassword, err := domain.HashPassword(password)
	if err != nil {
		return err
	}

	// Guardar en la base de datos con la contraseña encriptada
	return cp.db.Save(nombre, telefono, hashedPassword)
}
