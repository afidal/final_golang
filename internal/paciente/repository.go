package paciente

import (
	"errors"
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
	"tp_final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	Create(paciente dto.Paciente) (domain.Paciente, error)
	Update(id int, paciente dto.Paciente) error
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewPacienteRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {

	paciente, err := r.storage.ReadPaciente(id)
	if err != nil {
		return domain.Paciente{}, errors.New("No se ha encontrado al paciente solicitado")
	}
	return paciente, nil

}

func (r *repository) Create(paciente dto.Paciente) (domain.Paciente, error) {

	if r.storage.DniExists(paciente.Dni) {
		return domain.Paciente{}, errors.New("Ya hay un paciente registrado con el mismo DNI")
	}

	pacienteRetornado, err := r.storage.CreatePaciente(paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Se produjo un error cargando un nuevo paciente")
	}
	return pacienteRetornado, nil
}

func (r *repository) Update(id int, paciente dto.Paciente) error {

	p, err := r.storage.ReadPaciente(id)
	if err != nil {
		return errors.New("No se ha encontrado el paciente solicitado")
	}

	if p.Dni != paciente.Dni {
		if r.storage.DniExists(paciente.Dni){
			return errors.New("Ya hay un paciente registrado con el mismo DNI")
		}
	}

	err = r.storage.UpdatePaciente(id, paciente)
	if err != nil {
		return errors.New("Se produjo un error modificando el paciente solicitado")
	}

	return nil

}

func (r *repository) Delete(id int) error {

	err := r.storage.DeletePaciente(id)
	if err != nil {
		return err
	}

	return nil }
