package turno

import (
	"errors"
	"tp_final/internal/domain"
	"tp_final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Turno, error)
	Create(turno domain.Turno) (domain.Turno, error)
	Update(id int, turno domain.Turno) error
	Delete(id int) error
	//GetByDNI(dni string) ([]domain.Turno, error)
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

func (r *repository) Create(turno domain.Turno) (domain.Turno, error) {

	err := r.storage.ValidarOdontologoPacienteExist(turno)
	if err != nil {
		return domain.Turno{}, err
	}

	turno, err = r.storage.CreateTurno(turno)
	if err != nil {
		return domain.Turno{}, errors.New("Se produjo un error cargando un nuevo turno")
	}
	return turno, nil
}

func (r *repository) Update(id int, turno domain.Turno) error {

	_, err := r.storage.ReadTurnoId(id)
	if err != nil {
		return errors.New("No se ha encontrado el turno solicitado")
	}

	err = r.storage.ValidarOdontologoPacienteExist(turno)
	if err != nil {
		return err
	}

	err = r.storage.UpdateTurno(turno)
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
