package turno

import (
	"errors"
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
	"tp_final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Turno, error)
	Create(turno dto.Turno) (domain.Turno, error)
	Update(id int, turno dto.Turno) error
	Delete(id int) error
	GetByDNI(dni string) ([]dto.TurnoDatos, error)
	CreateDniMat(turno dto.TurnoAux) (dto.TurnoAuxId, error)
}

type repository struct {
	storage store.StoreInterface
}

func NewTurnoRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Turno, error) {

	turno, err := r.storage.ReadTurnoId(id)
	if err != nil {
		return domain.Turno{}, errors.New("No se ha encontrado al turno solicitado")
	}
	return turno, nil

}

func (r *repository) Create(turno dto.Turno) (domain.Turno, error) {

	err := r.storage.ValidarOdontologoPacienteExist(turno)
	if err != nil {
		return domain.Turno{}, err
	}

	turnoRetornado, err := r.storage.CreateTurno(turno)
	if err != nil {
		return domain.Turno{}, errors.New("Se produjo un error cargando un nuevo turno")
	}
	return turnoRetornado, nil
}

func (r *repository) Update(id int, turno dto.Turno) error {

	_, err := r.storage.ReadTurnoId(id)
	if err != nil {
		return errors.New("No se ha encontrado el turno solicitado")
	}

	err = r.storage.ValidarOdontologoPacienteExist(turno)
	if err != nil {
		return err
	}

	err = r.storage.UpdateTurno(id, turno)
	if err != nil {
		return errors.New("Se produjo un error modificando el turno solicitado")
	}

	return nil

}

func (r *repository) Delete(id int) error {

	err := r.storage.DeleteTurno(id)
	if err != nil {
		return err
	}

	return nil

}

func (r *repository) GetByDNI(dni string) ([]dto.TurnoDatos, error) {

	turnos, err := r.storage.ReadTurnoDni(dni)
	if err != nil {
		return []dto.TurnoDatos{}, errors.New("No se ha encontrado al turno solicitado")
	}
	return turnos, nil

}

func (r *repository) CreateDniMat(turno dto.TurnoAux) (dto.TurnoAuxId, error) {

	if !r.storage.MatriculaExists(turno.MatriculaOdontologo) {
		return dto.TurnoAuxId{}, errors.New("El odontólogo con la matrícula ingresada no está registrado en la base de datos")
	}

	if !r.storage.DniExists(turno.DniPaciente) {
		return dto.TurnoAuxId{}, errors.New("El paciente con el DNI ingresado no está registrado en la base de datos")
	}

	turnoRetornado, err := r.storage.CreateTurnoDniMat(turno)
	if err != nil {
		return dto.TurnoAuxId{}, errors.New("Se produjo un error cargando un nuevo turno")
	}
	return turnoRetornado, nil
}
