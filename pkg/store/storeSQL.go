package store

import (
	"database/sql"
	"errors"
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
	"fmt"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

// Odontologo

func (s *sqlStore) ReadOdontologo(id int) (domain.Odontologo, error) {

	var odontologo domain.Odontologo

	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	fmt.Println(*row)
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

func (s *sqlStore) CreateOdontologo(odontologo dto.Odontologo) (domain.Odontologo, error) {

	var odontologoRetornado domain.Odontologo

	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Odontologo{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	lastID, _ := result.LastInsertId()
	odontologoRetornado.Id = int(lastID)
	odontologoRetornado.Nombre = odontologo.Nombre
	odontologoRetornado.Apellido = odontologo.Apellido
	odontologoRetornado.Matricula = odontologo.Matricula

	return odontologoRetornado, nil

}

func (s *sqlStore) UpdateOdontologo(id int, odontologo dto.Odontologo) error {

	stmt, err := s.db.Prepare("UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *sqlStore) DeleteOdontologo(id int) error {

	stmt := "DELETE FROM odontologos WHERE id = ?;"

	result, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	var rows int64

	rows, err = result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("No existe un odontólogo con el ID indicado")
	}

	return nil
}

// Verifica si ya existe un odontólogo con la misma matrícula en la base de datos
func (s *sqlStore) MatriculaExists(matricula string) bool {

	var id int

	row := s.db.QueryRow("SELECT id FROM odontologos WHERE matricula = ?;", matricula)

	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}

// // Paciente

// func (s *sqlStore) ReadPaciente(id int) (domain.Paciente, error) {

// 	var paciente domain.Paciente

// 	query := "SELECT * FROM pacientes WHERE id = ?;"
// 	row := s.db.QueryRow(query, id)
// 	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaAlta)
// 	if err != nil {
// 		return domain.Paciente{}, err
// 	}

// 	return paciente, nil
// }

// func (s *sqlStore) CreatePaciente(paciente dto.Paciente) (domain.Paciente, error) {

// 	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?);"

// 	stmt, err := s.db.Prepare(query)
// 	if err != nil {
// 		return domain.Paciente{}, err
// 	}

// 	defer stmt.Close()

// 	result, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaAlta)
// 	if err != nil {
// 		return domain.Paciente{}, err
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return domain.Paciente{}, err
// 	}

// 	lastID, _ := result.LastInsertId()
// 	paciente.Id = int(lastID)

// 	return paciente, nil

// }

// func (s *sqlStore) UpdatePaciente(paciente dto.Paciente) error {

// 	stmt, err := s.db.Prepare("UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?;")
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaAlta, paciente.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *sqlStore) DeletePaciente(id int) error {

// 	stmt := "DELETE FROM pacientes WHERE id = ?;"

// 	result, err := s.db.Exec(stmt, id)
// 	if err != nil {
// 		return err
// 	}

// 	var rows int64

// 	rows, err = result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rows == 0 {
// 		return errors.New("No existe un paciente con el ID indicado")
// 	}

// 	return nil

// }

// // Verifica si ya existe un paciente con el mismo DNI en la base de datos
// // Asumimos que el DNI es único
// func (s *sqlStore) DniExists(dni string) bool {

// 	var id int

// 	row := s.db.QueryRow("SELECT id FROM pacientes WHERE dni = ?;", dni)

// 	err := row.Scan(&id)
// 	if err != nil {
// 		return false
// 	}

// 	if id > 0 {
// 		return true
// 	}

// 	return false
// }

// // Turno

// // func (s *sqlStore) ReadTurnoId(id int) (domain.Turno, error) {

// // 	var turno domain.Turno

// // 	query := "SELECT * FROM turnos WHERE id = ?;"
// // 	row := s.db.QueryRow(query, id)
// // 	fmt.Println("ID",*row)
// // 	err := row.Scan(&turno.Id, &turno.IdPaciente, &turno.IdOdontologo, &turno.Fecha, &turno.Hora, &turno.Descripcion)
// // 	if err != nil {
// // 		return domain.Turno{}, err
// // 	}

// // 	return turno, nil
// // }

// // Verifica si el paciente y el odontologo existen en la base de datos
// func (s *sqlStore) ValidarOdontologoPacienteExist(turno dto.Turno) error {

// 	_, err := s.ReadOdontologo(turno.IdOdontologo)
// 	if err != nil {
// 		return errors.New("El odontólogo al que se quiere asignar el turno no existe en la base de datos")
// 	}

// 	_, err = s.ReadPaciente(turno.IdPaciente)
// 	if err != nil {
// 		return errors.New("El paciente al que se quiere asignar el turno no existe en la base de datos")
// 	}

// 	return nil

// }

// func (s *sqlStore) CreateTurno(turno dto.Turno) (domain.Turno, error) {

// 	query := "INSERT INTO turnos (id_paciente, id_odontologo, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"

// 	stmt, err := s.db.Prepare(query)
// 	if err != nil {
// 		return domain.Turno{}, err
// 	}

// 	defer stmt.Close()

// 	result, err := stmt.Exec(turno.IdPaciente, turno.IdOdontologo, turno.Fecha, turno.Hora, turno.Descripcion)
// 	if err != nil {
// 		return domain.Turno{}, err
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return domain.Turno{}, err
// 	}

// 	lastID, _ := result.LastInsertId()
// 	turno.Id = int(lastID)

// 	return turno, nil

// }

// func (s *sqlStore) UpdateTurno(turno dto.Turno) error {

// 	stmt, err := s.db.Prepare("UPDATE turnos SET id_paciente = ?, id_odontologo = ?, fecha = ?, hora = ?, descripcion = ? WHERE id = ?;")
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(turno.IdPaciente, turno.IdOdontologo, turno.Fecha, turno.Hora, turno.Descripcion, turno.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func (s *sqlStore) DeleteTurno(id int) error {

// 	stmt := "DELETE FROM turnos WHERE id = ?;"

// 	result, err := s.db.Exec(stmt, id)
// 	if err != nil {
// 		return err
// 	}

// 	var rows int64

// 	rows, err = result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rows == 0 {
// 		return errors.New("No existe un turno con el ID indicado")
// 	}

// 	return nil
// }


// func (s *sqlStore) ReadTurnoDni(dni string) ([]dto.TurnoDatos, error) {

// 	var turnos []domain.TurnoDatos
// 	var turno domain.TurnoDatos
	
// 	query := "SELECT t.id, o.*, p.* , t.fecha, t.hora, t.descripcion FROM turnos t INNER JOIN odontologos o ON o.id = t.id_odontologo INNER JOIN pacientes p ON p.id = t.id_paciente WHERE p.dni = ?;"
// 	rows, err := s.db.Query(query, dni)
// 	if err != nil {
// 		return []domain.TurnoDatos{}, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
		
// 		err := rows.Scan(&turno.Id, &turno.Odontologo.Id, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula, &turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.Dni, &turno.Paciente.FechaAlta, &turno.Fecha, &turno.Hora, &turno.Descripcion)
// 		if err != nil {
// 			return []domain.TurnoDatos{}, err
// 		}

// 		turnos = append(turnos, turno)
// 	}

// 	return turnos, nil
// }


// func (s *sqlStore) CreateTurnoDniMat(taux dto.TurnoAux) (domain.TurnoAuxId, error) {

// 	// 1 - Buscamos a que id de paciente corresponde el dni
	
// 	var idPaciente int

// 	row := s.db.QueryRow("SELECT id FROM pacientes WHERE dni = ?;", taux.DniPaciente)

// 	err := row.Scan(&idPaciente)
// 	if err != nil {
// 		return domain.TurnoAux{}, err
// 	}

// 	fmt.Println("Id Paciente:", idPaciente)

// 	// 2 - Buscamos a que id de odontologo corresponde la matricula

// 	var idOdontologo int

// 	row = s.db.QueryRow("SELECT id FROM odontologos WHERE matricula = ?;", taux.MatriculaOdontologo)
	
// 	err = row.Scan(&idOdontologo)
// 	if err != nil {
// 		return domain.TurnoAux{}, err
// 	}
	
// 	fmt.Println("Id Odontologo:", idOdontologo)

// 	// Creamos el turno

// 	query := "INSERT INTO turnos (id_paciente, id_odontologo, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"

// 	stmt, err := s.db.Prepare(query)
// 	if err != nil {
// 		return domain.TurnoAux{}, err
// 	}

// 	defer stmt.Close()

// 	result, err := stmt.Exec(idPaciente, idOdontologo, taux.Fecha, taux.Hora, taux.Descripcion)
// 	if err != nil {
// 		return domain.TurnoAux{}, err
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return domain.TurnoAux{}, err
// 	}

// 	lastID, _ := result.LastInsertId()
// 	taux.Id = int(lastID)

// 	return taux, nil
// }