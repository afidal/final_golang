package domain

type TurnoAux struct {
	DniPaciente         string `json:"dni_paciente" binding:"required"`
	MatriculaOdontologo string `json:"matricula_odontologo" binding:"required"`
	Fecha               string `json:"fecha" binding:"required"`
	Hora                string `json:"hora" binding:"required"`
	Descripcion         string `json:"descripcion" binding:"required"`
}
