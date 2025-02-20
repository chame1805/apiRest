package applications

import "practica/src/Materias/domain"

type ListMateriasUseCase struct {
	db domain.IMateriaRepository
}


func NewListMateriasUseCase(db domain.IMateriaRepository) *ListMateriasUseCase {
	return &ListMateriasUseCase{db: db}
}

func (cp *ListMateriasUseCase) Execute() ([]map[string]interface{}, error) {
	return cp.db.GetList()
}