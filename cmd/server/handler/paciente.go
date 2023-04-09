package handler

import (
	"errors"
	"strconv"
	"tp_final/internal/domain"
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

func (h *pacienteHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {

		var paciente domain.Paciente

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

		pacienteCreado, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, pacienteCreado)

	}

}

func validarCamposPaciente(paciente *domain.Paciente) (bool, error) {

	if paciente.Nombre == "" || paciente.Apellido == "" || paciente.Domicilio == "" || paciente.Dni == "" || paciente.FechaAlta == "" {
		return false, errors.New("Ha ocurrido un error. Debe completar todos los campos")
	}

	return true, nil

}

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

		var paciente domain.Paciente

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

		err = h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, nil)

	}
}

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

		update := domain.Paciente{
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
