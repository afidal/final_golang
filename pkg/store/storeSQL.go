package store

import (
	"database/sql"
	"errors"
	"tp_final/internal/domain"
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
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
	if err != nil {
            return domain.Odontologo{}, err
        }

    return odontologo, nil
}

func (s *sqlStore) CreateOdontologo(odontologo domain.Odontologo) (domain.Odontologo, error) {

	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Odontologo{}, err
	}

	result, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	lastID, _ := result.LastInsertId()
	odontologo.Id = int(lastID)

	return odontologo, nil

}

func (s *sqlStore) UpdateOdontologo(odontologo domain.Odontologo) error {

	stmt, err := s.db.Prepare("UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;") 
    if err != nil {
        return err
    }

    defer stmt.Close()    

    _, err = stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula, odontologo.Id) 
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

// Paciente

func (s *sqlStore) ReadPaciente(id int) (domain.Paciente, error) {

	var paciente domain.Paciente

	query := "SELECT * FROM pacientes WHERE id = ?;"
    row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaAlta)
	if err != nil {
            return domain.Paciente{}, err
        }

    return paciente, nil
}

func (s *sqlStore) CreatePaciente(paciente domain.Paciente) (domain.Paciente, error) {

	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Paciente{}, err
	}

	result, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Paciente{}, err
	}

	lastID, _ := result.LastInsertId()
	paciente.Id = int(lastID)

	return paciente, nil

}

func (s *sqlStore) UpdatePaciente(paciente domain.Paciente) error {

	stmt, err := s.db.Prepare("UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?;") 
    if err != nil {
        return err
    }

    defer stmt.Close()    

    _, err = stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaAlta, paciente.Id) 
    if err != nil {
        return err
    }

    return nil
}

func (s *sqlStore) DeletePaciente(id int) error {

	stmt := "DELETE FROM pacientes WHERE id = ?;"

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
		return errors.New("No existe un paciente con el ID indicado")
	}

	return nil
}

// Turno

func (s *sqlStore) ReadTurnoId(id int) (domain.Turno, error) {
	return domain.Turno{}, nil
}

func (s *sqlStore) CreateTurno(turno domain.Turno) error {
	return nil
}

func (s *sqlStore) UpdateTurno(turno domain.Turno) error {
	return nil
}

func (s *sqlStore) DeleteTurno(id int) error {
	return nil
}

func (s *sqlStore) CreateTurnoDniMat(turno domain.TurnoAux) error {
	return nil
}

func (s *sqlStore) ReadTurnoDni(dni string) (domain.Turno, error) {
	return domain.Turno{}, nil
}
