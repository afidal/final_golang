package odontologo

import (
	"errors"
	"tp_final/internal/domain"
	"tp_final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewOdontologoRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Odontologo, error) {

	odontologo, err := r.storage.ReadOdontologo(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("No se ha encontrado al odontólogo solicitado")
	}
	return odontologo, nil

}

func (r *repository) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {

	if r.storage.MatriculaExists(odontologo.Matricula) {
		return domain.Odontologo{}, errors.New("Ya hay un odontólogo registrado con la misma matrícula")
	}

	odontologo, err := r.storage.CreateOdontologo(odontologo)
	if err != nil {
		return domain.Odontologo{}, errors.New("Se produjo un error cargando un nuevo odontólogo")
	}
	return odontologo, nil
}

func (r *repository) Update(id int, odontologo domain.Odontologo) error {

	od, err := r.storage.ReadOdontologo(id)
	if err != nil {
		return errors.New("No se ha encontrado al odontólogo solicitado")
	}

	if od.Matricula != odontologo.Matricula {
		if r.storage.MatriculaExists(odontologo.Matricula){
			return errors.New("Ya hay un odontólogo registrado con la misma matrícula")
		}
	}

	err = r.storage.UpdateOdontologo(odontologo)
	if err != nil {
		return errors.New("Se produjo un error modificando al odontólogo solicitado")
	}

	return nil

}

func (r *repository) Delete(id int) error {

	err := r.storage.DeleteOdontologo(id)
	if err != nil {
		return err
	}

	return nil

}
