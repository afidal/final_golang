package dto

type Turno struct {

	IdPaciente   int    `json:"id_paciente" binding:"required"`
	IdOdontologo int    `json:"id_odontologo" binding:"required"`
	Fecha        string `json:"fecha" binding:"required"`
	Hora         string `json:"hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
	
}