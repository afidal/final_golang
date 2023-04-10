package paciente

import (
	"tp_final/internal/domain"
)

type Service interface {
	
	GetByID(id int) (domain.Paciente, error)
	Create(paciente domain.Paciente) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) (error)
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

func (s *service) Create(paciente domain.Paciente) (domain.Paciente, error) {

	paciente, err := s.r.Create(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil

}

func (s *service) Update(id int, paciente domain.Paciente) error {

	p, err := s.r.GetByID(id)
	if err != nil {
		return err
	}

	if paciente.Nombre != "" {
		p.Nombre = paciente.Nombre
	}

	if paciente.Apellido != "" {
		p.Apellido = paciente.Apellido
	}

	if paciente.Domicilio != "" {
		p.Domicilio = paciente.Domicilio
	}

	if paciente.Dni != "" {
		p.Dni = paciente.Dni
	}

	if paciente.FechaAlta != "" {
		p.FechaAlta = paciente.FechaAlta
	}

	err = s.r.Update(id, p)
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