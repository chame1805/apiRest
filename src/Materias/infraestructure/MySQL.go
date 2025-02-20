package infraestructure

import (
	"fmt"
	"log"
	"practica/src/core"
)

type MySQLMaterias struct {
	conn *core.Conn_MySQL
}



func NewMySQLMaterias() *MySQLMaterias {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLMaterias{conn: conn}
}

func (mysql *MySQLMaterias) Save(name string) error {
	query := "INSERT INTO materias (name) VALUES (?)"
	log.Printf("Ejecutando consulta: %s, con valores: name=%s", query, name)
	result, err := mysql.conn.ExecutePreparedQuery(query, name)
	if err != nil {
		log.Printf("Error al guardar la materia: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se insertó ningún registro")
		return fmt.Errorf("no se insertó ningún registro")
	}
	log.Println("[MySQL] - Materia guardada exitosamente")
	return nil
}

func (mysql *MySQLMaterias) GetList() ([]map[string]interface{}, error) {
	query := "SELECT * FROM materias"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	var materias []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		materias = append(materias, map[string]interface{}{"id": id, "name": name})
	}
	return materias, nil
}

func (mysql *MySQLMaterias) GetMateria(id int) (map[string]interface{}, error) {
	query := "SELECT * FROM materias WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()
	var materia map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		materia = map[string]interface{}{"id": id, "name": name}
	}
	return materia, nil
}

func (mysql *MySQLMaterias) Update(id int, name string) error {
	query := "UPDATE materias SET name = ? WHERE id = ?"
	log.Printf("Ejecutando consulta: %s, con valores: name=%s, id=%d", query, name, id)
	result, err := mysql.conn.ExecutePreparedQuery(query, name, id)
	if err != nil {
		log.Printf("Error al actualizar la materia: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se actualizó ningún registro")
		return fmt.Errorf("no se actualizó ningún registro")
	}
	log.Println("[MySQL] - Materia actualizada exitosamente")
	return nil
}

func (mysql *MySQLMaterias) Delete(id int) error {
	query := "DELETE FROM materias WHERE id = ?"
	log.Printf("Ejecutando consulta: %s, con valores: id=%d", query, id)
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar la materia: %v", err)
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se eliminó ningún registro")
		return fmt.Errorf("no se eliminó ningún registro")
	}
	log.Println("[MySQL] - Materia eliminada exitosamente")
	return nil
}
