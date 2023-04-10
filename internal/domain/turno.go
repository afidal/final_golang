package domain

type Turno struct {
	Id           int    `json:"id"`
	IdPaciente   int    `json:"id_paciente" binding:"required"`
	IdOdontologo int    `json:"id_odontologo" binding:"required"`
	Fecha        string `json:"fecha" binding:"required"`
	Hora         string `json:"hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}

type TurnoDatos struct {
	Id          int        `json:"id"`
	Odontologo  Odontologo `json:"odontologo" binding:"required"`
	Paciente    Paciente   `json:"paciente" binding:"required"`
	Fecha       string     `json:"fecha" binding:"required"`
	Hora        string     `json:"hora" binding:"required"`
	Descripcion string     `json:"descripcion" binding:"required"`
}

type TurnoAux struct {
	Id                  int    `json:"id"`
	DniPaciente         string `json:"dni_paciente" binding:"required"`
	MatriculaOdontologo string `json:"matricula_odontologo" binding:"required"`
	Fecha               string `json:"fecha" binding:"required"`
	Hora                string `json:"hora" binding:"required"`
	Descripcion         string `json:"descripcion" binding:"required"`
}
