package domain

type Turno struct {
	Id           int    `json:"id"`
	IdPaciente   int    `json:"id_paciente" binding:"required"`
	IdOdontologo int    `json:"id_odontologo" binding:"required"`
	Fecha        string `json:"fecha" binding:"required"`
	Hora         string `json:"hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}
