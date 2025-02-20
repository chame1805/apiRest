package applications

import (
	"practica/src/Materias/domain"
)

type CreateMateriaUseCase struct {
	db domain.IMateriaRepository
}

func NewCreateMateriaUseCase(db domain.IMateriaRepository) *CreateMateriaUseCase {
	return &CreateMateriaUseCase{db: db}
}
//este metodo llama al metodo save para guardar el nombre de la materia
func (c *CreateMateriaUseCase) Execute(name string) {
	c.db.Save(name)
}
