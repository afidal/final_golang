package handler

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"tp_final/internal/domain/dto"
	"tp_final/internal/paciente"
	"tp_final/pkg/web"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// GetByID godoc
// @Summary      GET paciente by ID
// @Description  Obtiene un paciente por su ID
// @Tags         Paciente
// @Produce      json
// @Param        id path int true "Paciente Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Router       /pacientes/:id [get]
func (h *pacienteHandler) GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al paciente solicitado"))
			return
		}
		web.Success(c, 200, paciente)
	}

}

// Post godoc
// @Summary      POST paciente
// @Description  Crea un nuevo paciente
// @Tags         Paciente
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Success      201 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Router       /pacientes [post]
func (h *pacienteHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {

		var paciente dto.Paciente

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposPaciente(&paciente)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		paciente.Dni = strings.ReplaceAll(paciente.Dni, ".", "")

		fechaValida, err := validarFechaAlta(paciente.FechaAlta)
		if !fechaValida {
			web.Failure(c, 400, err)
			return
		}

		pacienteCreado, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, pacienteCreado)

	}

}

// Put godoc
// @Summary      PUT paciente by ID
// @Description  Actualiza un paciente por su ID
// @Tags         Paciente
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body dto.Paciente true "Paciente"
// @Param        id path int true "Odontologo Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /pacientes/:id [put]
func (h *pacienteHandler) Put() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al paciente solicitado"))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		var paciente dto.Paciente

		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposPaciente(&paciente)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		paciente.Dni = strings.ReplaceAll(paciente.Dni, ".", "")

		fechaValida, err := validarFechaAlta(paciente.FechaAlta)
		if !fechaValida {
			web.Failure(c, 400, err)
			return
		}

		err = h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, nil)

	}
}

// Patch godoc
// @Summary      PATCH paciente by ID
// @Description  Actualizar parcialmente un paciente por su ID
// @Tags         Paciente
// @Accept		 json
// @Produce      json
// @Param        token header string true "token"
// @Param        body body dto.Paciente true "Paciente"
// @Param        id path int true "Paciente Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /pacientes/:id [patch]
func (h *pacienteHandler) Patch() gin.HandlerFunc {

	type PatchRequest struct {
		Nombre    string `json:"nombre,omitempty"`
		Apellido  string `json:"apellido,omitempty"`
		Domicilio string `json:"domicilio,omitempty"`
		Dni       string `json:"dni,omitempty"`
		FechaAlta string `json:"fecha_alta,omitempty"`
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
			web.Failure(c, 404, errors.New("No se ha encontrado al paciente solicitado"))
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

		if request.FechaAlta != "" {
			fechaValida, err := validarFechaAlta(request.FechaAlta)
			if !fechaValida {
				web.Failure(c, 400, err)
				return
			}
		}

		request.Dni = strings.ReplaceAll(request.Dni, ".", "")

		update := dto.Paciente{
			Nombre:    request.Nombre,
			Apellido:  request.Apellido,
			Domicilio: request.Domicilio,
			Dni:       request.Dni,
			FechaAlta: request.FechaAlta,
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
// @Summary      DELETE paciente by ID
// @Description  Elimina un paciente por su ID
// @Tags         Paciente
// @Produce      json
// @Param        token header string true "token"
// @Param        id path int true "Odontologo Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
// @Router       /pacientes/:id [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {

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

// Fx para validaciones de datos

func validarCamposPaciente(paciente *dto.Paciente) (bool, error) {

	if paciente.Nombre == "" || paciente.Apellido == "" || paciente.Domicilio == "" || paciente.Dni == "" || paciente.FechaAlta == "" {
		return false, errors.New("Ha ocurrido un error. Debe completar todos los campos")
	}

	return true, nil

}

func validarFechaAlta(fecha string) (bool, error) {
	re := regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[/](0[1-9]|1[012])[/](19|20)\d\d`)
	if !re.MatchString(fecha) {
		return false, errors.New("La fecha ingresada es inválida. Debe tener el formato: dd/mm/yyyy")
	}

	return true, nil
}
