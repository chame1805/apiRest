package domain

type IMateriaRepository interface {
	Save(name string) error
	GetMateria(id int) (map[string]interface{}, error)
	GetList() ([]map[string]interface{}, error)
	Delete(id int) error
	Update(id int, name string) error
}