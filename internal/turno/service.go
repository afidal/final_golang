package turno

import (
	"tp_final/internal/domain"
)

type Service interface {
	
	GetByID(id int) (domain.Turno, error)
	Create(turno domain.Turno) (domain.Turno, error)
	Update(id int, turno domain.Turno) (error)
	Delete(id int) error
	GetByDNI(dni string) ([]domain.TurnoDatos, error)

}

type service struct {
	r Repository
}

func NewTurnoService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Turno, error) {

	turno, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *service) Create(turno domain.Turno) (domain.Turno, error) {

	t, err := s.r.Create(turno)
	if err != nil {
		return domain.Turno{}, err
	}

	return t, nil

}

func (s *service) Update(id int, turno domain.Turno) error {

	t, err := s.r.GetByID(id)
	if err != nil {
		return err
	}

	if turno.IdPaciente != 0 {
		t.IdPaciente = turno.IdPaciente
	}

	if turno.IdOdontologo != 0 {
		t.IdOdontologo = turno.IdOdontologo
	}

	if turno.Fecha != "" {
		t.Fecha = turno.Fecha
	}

	if turno.Hora != "" {
		t.Hora = turno.Hora
	}

	if turno.Descripcion != "" {
		t.Descripcion = turno.Descripcion
	}
	
	err = s.r.Update(id, t)
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

func (s *service) GetByDNI(dni string) ([]domain.TurnoDatos, error) {

	turnos, err := s.r.GetByDNI(dni)
	if err != nil {
		return []domain.TurnoDatos{}, err
	}
	return turnos, nil

}