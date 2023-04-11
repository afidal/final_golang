package store

import (
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
)

type StoreInterface interface {

	// Odontologo

	ReadOdontologo(id int) (domain.Odontologo, error)

	CreateOdontologo(odontologo dto.Odontologo) (domain.Odontologo, error)

	UpdateOdontologo(id int, odontologo dto.Odontologo) error

	DeleteOdontologo(id int) error

	MatriculaExists(matricula string) bool

	// Paciente

	ReadPaciente(id int) (domain.Paciente, error)

	CreatePaciente(paciente dto.Paciente) (domain.Paciente, error)

	UpdatePaciente(id int, paciente dto.Paciente) error

	DeletePaciente(id int) error

	DniExists(dni string) bool

	// // Turno

	ReadTurnoId(id int) (domain.Turno, error)

	CreateTurno(turno dto.Turno) (domain.Turno, error)

	UpdateTurno(id int, turno dto.Turno) error

	DeleteTurno(id int) error

	ReadTurnoDni(dni string) ([]dto.TurnoDatos, error)

	CreateTurnoDniMat(turno dto.TurnoAux) (dto.TurnoAuxId, error)

	ValidarOdontologoPacienteExist(turno dto.Turno) error
}
