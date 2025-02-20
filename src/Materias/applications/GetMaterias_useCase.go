package applications

import "practica/src/Materias/domain"

type GetMateriaUseCase struct {
	db domain.IMateriaRepository
}

func NewGetMateriaUseCase(db domain.IMateriaRepository) *GetMateriaUseCase {
	return &GetMateriaUseCase{db: db}
}

func (cp *GetMateriaUseCase) Execute(id int) (map[string]interface{}, error) {
	return cp.db.GetMateria(id)
}