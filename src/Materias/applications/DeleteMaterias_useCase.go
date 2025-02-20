package applications

import(

	"practica/src/Materias/domain"
)

type DeleteMateriaUseCase struct {
	db domain.IMateriaRepository
}

func NewDeleteMateriaUseCase(db domain.IMateriaRepository) *DeleteMateriaUseCase {
	return &DeleteMateriaUseCase{db: db}
}

func (c *DeleteMateriaUseCase) Execute(id int) error {
	return c.db.Delete(id)
}