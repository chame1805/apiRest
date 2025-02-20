package applications

import "practica/src/Materias/domain"

type UpdateMateriaUseCase struct {
	db domain.IMateriaRepository
}

func NewUpdateMateriaUseCase(db domain.IMateriaRepository) *UpdateMateriaUseCase {
	return &UpdateMateriaUseCase{db: db}
}

func (cp *UpdateMateriaUseCase) Execute(id int, name string) error {
	return cp.db.Update(id, name)
}