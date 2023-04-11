package paciente

import (
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
)

type Service interface {
	
	GetByID(id int) (domain.Paciente, error)
	Create(paciente dto.Paciente) (domain.Paciente, error)
	Update(id int, paciente dto.Paciente) (error)
	Delete(id int) error
	
}

type service struct {
	r Repository
}

func NewPacienteService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Paciente, error) {

	paciente, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *service) Create(paciente dto.Paciente) (domain.Paciente, error) {

	pacienteRetornado, err := s.r.Create(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}

	return pacienteRetornado, nil

}

func (s *service) Update(id int, p dto.Paciente) error {

	paciente, err := s.r.GetByID(id)
	if err != nil {
		return err
	}

	if p.Nombre != "" {
		paciente.Nombre = p.Nombre
	}

	if p.Apellido != "" {
		paciente.Apellido = p.Apellido
	}

	if p.Domicilio != "" {
		paciente.Domicilio = p.Domicilio
	}

	if p.Dni != "" {
		paciente.Dni = p.Dni
	}

	if p.FechaAlta != "" {
		paciente.FechaAlta = p.FechaAlta
	}

	var pacienteActualizado dto.Paciente
	
	pacienteActualizado.Nombre = paciente.Nombre
	pacienteActualizado.Apellido = paciente.Apellido
	pacienteActualizado.Domicilio = paciente.Domicilio
	pacienteActualizado.Dni = paciente.Dni
	pacienteActualizado.FechaAlta = paciente.FechaAlta
	

	err = s.r.Update(id, pacienteActualizado)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) Delete(id int) error {
	
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil

}