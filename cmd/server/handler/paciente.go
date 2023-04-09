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

// GetByID godoc
// @Summary      GET paciente by ID
// @Description  Obtiene un paciente por su ID
// @Tags         domain.Paciente
// @Produce      json
// @Param        id path int true "Paciente Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
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
// @Tags         domain.Paciente
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Success      201 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Router       /pacientes [post]
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

// func validateFormatoFecha(fecha string) (bool, error) {

// 	partes := strings.Split(fecha, "/")

// 	partes_number := []int{}

// 	if len(partes) != 3 {
// 		return false, errors.New("La fecha debe tener el formato dd/mm/yyyy")
// 	}



// 	for parte := range partes {
// 		number, err := strconv.Atoi(partes[parte])
// 		if err != nil {
// 			return false, errors.New("Fecha inválida, debe contener solo números")
// 		}
// 		partes_number = append(partes_number, number)
// 	}

// 	condition := (partes_number[0] < 1 || partes_number[0] > 31) && (partes_number[1] < 1 || partes_number[1] > 12) && (list[2] < 1 || list[2] > 9999)

// 	if condition {
// 		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
// 	}

// 	return true, nil

// }

// func validateDni(dni string) (string, error) {

// 	valid_dni := strings.ReplaceAll(dni, ".", "")

// 	return valid_dni, nil

// }


// Put godoc
// @Summary      PUT paciente by ID
// @Description  Actualiza un paciente por su ID
// @Tags         domain.Paciente
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Param        id path int true "Odontologo Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
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

// Patch godoc
// @Summary      PATCH paciente by ID
// @Description  Actualizar parcialmente un paciente por su ID
// @Tags         domain.Paciente
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Param        id path int true "Paciente Id"
// @Success      200 {object} web.response
// @Failure      400 {object} web.errorResponse
// @Failure      404 {object} web.errorResponse
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

// Delete godoc
// @Summary      DELETE paciente by ID
// @Description  Elimina un paciente por su ID
// @Tags         domain.Paciente
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