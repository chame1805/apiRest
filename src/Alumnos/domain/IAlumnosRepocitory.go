package domain

type IAlumnoRepository interface {
	Save(nombre string, telefono string,password string) error
	GetAll() ([]map[string]interface{}, error)
	GetAlumno(id int) (map[string]interface{}, error)
	Update(id int, nombre string, telefono string) error
	Delete(id int) error
}
