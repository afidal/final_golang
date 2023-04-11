package dto

import (
	"tp_final/internal/domain"
)

type TurnoDatos struct {

	Id          int        `json:"id"`
	Odontologo  domain.Odontologo `json:"odontologo" binding:"required"`
	Paciente    domain.Paciente   `json:"paciente" binding:"required"`
	Fecha       string     `json:"fecha" binding:"required"`
	Hora        string     `json:"hora" binding:"required"`
	Descripcion string     `json:"descripcion" binding:"required"`
	
}