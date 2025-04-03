package infraestructure

import (
	"fmt"
	"log"
	"practica/src/core"
	//"practica/src/Alumnos/domain"
)

type MySQLAlumnos struct {
	conn *core.Conn_MySQL
}

// Constructor
func NewMySQLAlumnos() *MySQLAlumnos {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLAlumnos{conn: conn}
}

// Implementación de Save con error
func (mysql *MySQLAlumnos) Save(nombre string, telefono string,password string) error {
	query := "INSERT INTO alumnos (nombre, telefono,password) VALUES (?, ?,?)"

	// Agregar logs para depuración
	log.Printf("Ejecutando consulta: %s, con valores: nombre=%s, telefono=%s password=%s", query, nombre, telefono,password)

	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, telefono,password)
	if err != nil {
		log.Printf("Error al guardar el alumno: %v", err)
		return err
	}

	// Verificar si realmente se insertó un registro
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se insertó ningún registro")
		return fmt.Errorf("no se insertó ningún registro")
	}

	log.Println("[MySQL] - Alumno guardado exitosamente")
	return nil
}

// Implementación de GetAll con error
func (mysql *MySQLAlumnos) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT * FROM alumnos"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var alumnos []map[string]interface{}
	for rows.Next() {
		var id int
		var nombre string
		var telefono string
		if err := rows.Scan(&id, &nombre, &telefono); err != nil {
			return nil, err
		}
		alumno := map[string]interface{}{
			"id":       id,
			"nombre":   nombre,
			"telefono": telefono,
		}
		alumnos = append(alumnos, alumno)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return alumnos, nil
}

// Implementación de GetAlumno con error
func (mysql *MySQLAlumnos) GetAlumno(id int) (map[string]interface{}, error) {
	query := "SELECT * FROM alumnos WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var alumno map[string]interface{}
	if rows.Next() {
		var id int
		var nombre string
		var telefono string
		if err := rows.Scan(&id, &nombre, &telefono); err != nil {
			return nil, err
		}
		alumno = map[string]interface{}{
			"id":       id,
			"nombre":   nombre,
			"telefono": telefono,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return alumno, nil
}

// Implementación de Update con error
func (mysql *MySQLAlumnos) Update(id int, nombre string, telefono string) error {
	query := "UPDATE alumnos SET nombre = ?, telefono = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, telefono, id)
	if err != nil {
		log.Printf("Error al actualizar el alumno: %v", err)
		return err
	}

	log.Println("[MySQL] - Alumno actualizado exitosamente")
	return nil
}

// Implementación de Delete con error
func (mysql *MySQLAlumnos) Delete(id int) error {
	query := "DELETE FROM alumnos WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar el alumno: %v", err)
		return err
	}

	log.Println("[MySQL] - Alumno eliminado exitosamente")
	return nil
}
