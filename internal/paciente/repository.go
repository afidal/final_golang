package paciente

import (
	"errors"
	"tp_final/internal/domain"
	"tp_final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	Create(odontologo domain.Paciente) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) error
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

func (r *repository) Create(paciente domain.Paciente) (domain.Paciente, error) {

	// if r.storage.MatriculaExists(odontologo.Matricula) {
	// 	return domain.Odontologo{}, errors.New("Ya hay un odontólogo registrado con la misma matrícula")
	// }

	paciente, err := r.storage.CreatePaciente(paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Se produjo un error cargando un nuevo paciente")
	}
	return paciente, nil
}

func (r *repository) Update(id int, paciente domain.Paciente) error {

	_, err := r.storage.ReadPaciente(id)
	if err != nil {
		return errors.New("No se ha encontrado al paciente solicitado")
	}

	// if pa.Matricula != odontologo.Matricula {
	// 	if r.storage.MatriculaExists(odontologo.Matricula){
	// 		return errors.New("Ya hay un odontólogo registrado con la misma matrícula")
	// 	}
	// }

	err = r.storage.UpdatePaciente(paciente)
	if err != nil {
		return errors.New("Se produjo un error modificando al paciente solicitado")
	}

	return nil

}

func (r *repository) Delete(id int) error {

	err := r.storage.DeletePaciente(id)
	if err != nil {
		return err
	}

	return nil

}
