package turno

import (
	"tp_final/internal/domain"
	"tp_final/internal/domain/dto"
)

type Service interface {
	GetByID(id int) (domain.Turno, error)
	Create(turno dto.Turno) (domain.Turno, error)
	Update(id int, turno dto.Turno) error
	Delete(id int) error
	GetByDNI(dni string) ([]dto.TurnoDatos, error)
	CreateDniMat(turno dto.TurnoAux) (dto.TurnoAuxId, error)
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

func (s *service) Create(turno dto.Turno) (domain.Turno, error) {

	turnoRetornado, err := s.r.Create(turno)
	if err != nil {
		return domain.Turno{}, err
	}

	return turnoRetornado, nil

}

func (s *service) Update(id int, t dto.Turno) error {

	turno, err := s.r.GetByID(id)
	if err != nil {
		return err
	}

	if t.IdPaciente != 0 {
		turno.IdPaciente = t.IdPaciente
	}

	if t.IdOdontologo != 0 {
		turno.IdOdontologo = t.IdOdontologo
	}

	if t.Fecha != "" {
		turno.Fecha = t.Fecha
	}

	if t.Hora != "" {
		turno.Hora = t.Hora
	}

	if t.Descripcion != "" {
		turno.Descripcion = t.Descripcion
	}


	var turnoActualizado dto.Turno
	
	turnoActualizado.IdPaciente = turno.IdPaciente
	turnoActualizado.IdOdontologo = turno.IdOdontologo
	turnoActualizado.Fecha = turno.Fecha
	turnoActualizado.Hora = turno.Hora
	turnoActualizado.Descripcion = turno.Descripcion


	err = s.r.Update(id, turnoActualizado)
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

func (s *service) GetByDNI(dni string) ([]dto.TurnoDatos, error) {

	turnos, err := s.r.GetByDNI(dni)
	if err != nil {
		return []dto.TurnoDatos{}, err
	}
	return turnos, nil

}

func (s *service) CreateDniMat(turno dto.TurnoAux) (dto.TurnoAuxId, error) {

	t, err := s.r.CreateDniMat(turno)
	if err != nil {
		return dto.TurnoAuxId{}, err
	}

	return t, nil

}