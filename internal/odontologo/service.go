package odontologo

import (
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
)

type Service interface {
	
	GetByID(id int) (domain.Odontologo, error)
	Create(odontologo dto.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo dto.Odontologo) (error)
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

func (s *service) Create(odontologo dto.Odontologo) (domain.Odontologo, error) {

	var odontologoRetornado domain.Odontologo

	odontologoRetornado, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologoRetornado, nil

}

func (s *service) Update(id int, o dto.Odontologo) error {

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


	var odontologoActualizado dto.Odontologo
	
	odontologoActualizado.Nombre = odontologo.Nombre
	odontologoActualizado.Apellido = odontologo.Apellido
	odontologoActualizado.Matricula = odontologo.Matricula

	err = s.r.Update(id, odontologoActualizado)
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