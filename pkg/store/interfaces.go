package store

import (
	"tp_final/internal/domain"
)

type StoreInterface interface {

	// Odontologo

	ReadOdontologo(id int) (domain.Odontologo, error)

	CreateOdontologo(odontologo domain.Odontologo) (domain.Odontologo, error)

	UpdateOdontologo(odontologo domain.Odontologo) error

	DeleteOdontologo(id int) error

	MatriculaExists(matricula string) bool

	// Paciente
	
	ReadPaciente(id int) (domain.Paciente, error)

	CreatePaciente(paciente domain.Paciente) (domain.Paciente, error)

	UpdatePaciente(paciente domain.Paciente) error

	DeletePaciente(id int) error

	// Turno

	ReadTurnoId(id int) (domain.Turno, error)

	CreateTurno(turno domain.Turno) (domain.Turno, error)

	UpdateTurno(turno domain.Turno) error

	DeleteTurno(id int) error

	CreateTurnoDniMat(turno domain.TurnoAux) (domain.Turno, error)

	ReadTurnoDni(dni string) ([]domain.Turno, error)

}
