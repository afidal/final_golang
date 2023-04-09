package odontologo

import (
	"tp_final/internal/domain"
)

type Service interface {
	
	GetByID(id int) (domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) (error)
	Delete(id int) error

}

type service struct {
	r Repository
}

func NewOdontologoService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {

	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *service) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {

	odontologo, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil

}

func (s *service) Update(id int, o domain.Odontologo) error {

	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return err
	}

	if o.Nombre != "" {
		odontologo.Nombre = o.Nombre
	}

	if o.Apellido != "" {
		odontologo.Apellido = o.Apellido
	}

	if o.Matricula != "" {
		odontologo.Matricula = o.Matricula
	}

	err = s.r.Update(id, odontologo)
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