package handler

import (
	"errors"
	"regexp"
	"strconv"
	"tp_final/internal/domain/dto"
	"tp_final/internal/turno"
	"tp_final/pkg/web"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// GetByID godoc
// @Summary      GET turno by ID
// @Description  Obtiene un turno por su ID
// @Tags         Turno
// @Produce      json
// @Param        id path int true "Turno Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Router       /turnos/:id [get]
func (h *turnoHandler) GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al turno solicitado"))
			return
		}
		web.Success(c, 200, turno)
	}

}

// Post godoc
// @Summary      POST turno
// @Description  Crea un nuevo turno con el ID del paciente y el ID del odontólogo
// @Tags         Turno
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Turno true "Turno"
// @Success      201 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Router       /turnos [post]
func (h *turnoHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {

		var turno dto.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposTurno(&turno)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		fechaValida, err := validarFecha(turno.Fecha)
		if !fechaValida {
			web.Failure(c, 400, err)
			return
		}

		horaValida, err := validarHora(turno.Hora)
		if !horaValida {
			web.Failure(c, 400, err)
			return
		}

		turnoCreado, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, turnoCreado)

	}

}

// Put godoc
// @Summary      PUT turno by ID
// @Description  Actualiza un turno por su ID
// @Tags         Turno
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Turno true "Turno"
// @Param        id path int true "Turno Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /turnos/:id [put]
func (h *turnoHandler) Put() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al turno solicitado"))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		var turno dto.Turno

		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposTurno(&turno)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		fechaValida, err := validarFecha(turno.Fecha)
		if !fechaValida {
			web.Failure(c, 400, err)
			return
		}

		horaValida, err := validarHora(turno.Hora)
		if !horaValida {
			web.Failure(c, 400, err)
			return
		}

		err = h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, nil)

	}
}

// Patch godoc
// @Summary      PATCH turno by ID
// @Description  Actualizar parcialmente un turno por su ID
// @Tags         Turno
// @Produce      json
// @Param        token header string true "token"
// @Param        body body dto.Turno true "Turno"
// @Param        id path int true "Turno Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /turnos/:id [patch]
func (h *turnoHandler) Patch() gin.HandlerFunc {

	type PatchRequest struct {
		IdPaciente   int    `json:"id_paciente,omitempty"`
		IdOdontologo int    `json:"id_odontologo,omitempty"`
		Fecha        string `json:"fecha,omitempty"`
		Hora         string `json:"hora,omitempty"`
		Descripcion  string `json:"descripcion,omitempty"`
	}

	return func(c *gin.Context) {

		var request PatchRequest

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el turno solicitado"))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		err = c.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		if request.Fecha != "" {
			fechaValida, err := validarFecha(request.Fecha)
			if !fechaValida {
				web.Failure(c, 400, err)
				return
			}
		}

		if request.Hora != "" {
			horaValida, err := validarHora(request.Hora)
			if !horaValida {
				web.Failure(c, 400, err)
				return
			}
		}

		update := dto.Turno{
			IdPaciente:   request.IdPaciente,
			IdOdontologo: request.IdOdontologo,
			Fecha:        request.Fecha,
			Hora:         request.Hora,
			Descripcion:  request.Descripcion,
		}

		err = h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, nil)

	}
}

// Delete godoc
// @Summary      DELETE turno by ID
// @Description  Elimina un turno por su ID
// @Tags         Turno
// @Produce      json
// @Param        token header string true "token"
// @Param        id path int true "Turno Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Router       /turnos/:id [delete]
func (h *turnoHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil)
	}
}

// GetByDNI godoc
// @Summary      GET turnos by DNI
// @Description  Obtiene una lista de turnos por DNI
// @Tags         Turno
// @Produce      json
// @Param        dni query string true "Turno Dni"
// @Success      200 {object} web.response
// @Failure      404 {object} web.errorResponse
// @Router       /turnos [get]
func (h *turnoHandler) GetByDNI() gin.HandlerFunc {

	return func(c *gin.Context) {

		dniParam := c.Query("dni")

		turnos, err := h.s.GetByDNI(dniParam)
		if err != nil {
			web.Failure(c, 404, errors.New("No se han encontrado turnos con el DNI ingresado"))
			return
		}
		web.Success(c, 200, turnos)
	}

}

// Post godoc
// @Summary      POST turno con DNI y Matrícula
// @Description  Crea un nuevo turno con el DNI del paciente y la matrícula del odontólogo
// @Tags         Turno
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body dto.TurnoAux true "TurnoAux"
// @Success      201 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Router       /turnos/DniMat [post]
func (h *turnoHandler) PostDniMat() gin.HandlerFunc {

	return func(c *gin.Context) {

		var turno dto.TurnoAux

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposTurnoAux(&turno)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		fechaValida, err := validarFecha(turno.Fecha)
		if !fechaValida {
			web.Failure(c, 400, err)
			return
		}

		horaValida, err := validarHora(turno.Hora)
		if !horaValida {
			web.Failure(c, 400, err)
			return
		}

		turnoCreado, err := h.s.CreateDniMat(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, turnoCreado)

	}

}

// Fx para validaciones de datos

func validarCamposTurnoAux(turno *dto.TurnoAux) (bool, error) {

	if turno.DniPaciente == "" || turno.MatriculaOdontologo == "" || turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "" {
		return false, errors.New("Ha ocurrido un error. Debe completar todos los campos")
	}

	return true, nil

}

func validarFecha(fecha string) (bool, error) {

	re := regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[/](0[1-9]|1[012])[/](19|20)\d\d`)
	if !re.MatchString(fecha) {
		return false, errors.New("La fecha ingresada es inválida. Debe tener el formato: dd/mm/yyyy")
	}

	return true, nil
}

func validarHora(hora string) (bool, error) {

	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)
	if !re.MatchString(hora) {
		return false, errors.New("La hora ingresada es inválida. Debe tener el formato: hh:mm")
	}

	return true, nil

}

func validarCamposTurno(turno *dto.Turno) (bool, error) {

	if turno.IdPaciente == 0 || turno.IdOdontologo == 0 || turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "" {
		return false, errors.New("Ha ocurrido un error. Debe completar todos los campos")
	}

	return true, nil

}
