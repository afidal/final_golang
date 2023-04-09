package handler

import (
	"errors"
	"strconv"
	"tp_final/internal/domain"
	"tp_final/internal/odontologo"
	"tp_final/pkg/web"

	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// GetByID godoc
// @Summary      GET odontologo by ID
// @Description  Obtiene un odontólogo por su ID
// @Tags         domain.Odontologo
// @Produce      json
// @Param        id path int true "Odontologo Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/:id [get]
func (h *odontologoHandler) GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al odontólogo solicitado"))
			return
		}
		web.Success(c, 200, odontologo)
	}

}

// Post godoc
// @Summary      POST odontologo
// @Description  Crea un nuevo odontólogo
// @Tags         domain.Odontologo
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /odontologos [post]
func (h *odontologoHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {

		var odontologo domain.Odontologo

		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposOdontologo(&odontologo)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		odontologoCreado, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, odontologoCreado)

	}

}

func validarCamposOdontologo(odontologo *domain.Odontologo) (bool, error) {

	if odontologo.Nombre == "" || odontologo.Apellido == "" || odontologo.Matricula == "" {
		return false, errors.New("Ha ocurrido un error. Debe completar todos los campos")
	}

	return true, nil

}


// Put godoc
// @Summary      PUT odontologo by ID
// @Description  Actualiza un odontologo por su ID
// @Tags         domain.Odontologo
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Odontologo true "Odontologo"
// @Param        id   path      int  true  "Odontologo Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /odontologos/:id [put]
func (h *odontologoHandler) Put() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El ID es inválido"))
			return
		}

		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado al odontologo solicitado"))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		var odontologo domain.Odontologo

		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("Json inválido"))
			return
		}

		camposValidos, err := validarCamposOdontologo(&odontologo)
		if !camposValidos {
			web.Failure(c, 400, err)
			return
		}

		err = h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, nil)

	}
}

// Patch godoc
// @Summary      PATCH odontologo by ID
// @Description  Actualizar parcialmente un odontólogo por su ID
// @Tags         domain.Odontologo
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Odontologo true "Odontologo"
// @Param        id   path      int  true  "Odontologo Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object} web.errorResponse
// @Router       /odontologos/:id [patch]
func (h *odontologoHandler) Patch() gin.HandlerFunc {

	type PatchRequest struct {
		Nombre    string `json:"nombre,omitempty"`
		Apellido  string `json:"apellido,omitempty"`
		Matricula string `json:"matricula,omitempty"`
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
			web.Failure(c, 404, errors.New("No se ha encontrado al odontologo solicitado"))
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

		update := domain.Odontologo{
			Nombre:    request.Nombre,
			Apellido:  request.Apellido,
			Matricula: request.Matricula,
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
// @Summary      DELETE odontologo by ID
// @Description  Elimina un odontólogo por su ID
// @Tags         domain.Odontologo
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Odontologo Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/:id [delete]
func (h *odontologoHandler) Delete() gin.HandlerFunc {

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
