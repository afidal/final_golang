package main

import (
	"database/sql"
	"os"
	"tp_final/cmd/server/handler"
	"tp_final/docs"
	"tp_final/internal/odontologo"
	"tp_final/internal/paciente"
	"tp_final/internal/turno"
	"tp_final/pkg/middleware"
	"tp_final/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API consultorio odontológico
// @version 1.0
// @description Esta API permite realizar operaciones CRUD sobre la base de datos del consultorio, que contiene registros de odontólogos, pacientes y turnos
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	db, err := sql.Open("mysql", "root:ek434603@tcp(localhost:3306)/clinica")
	if err != nil {
		panic(err.Error())
	}
	storage := store.NewSqlStore(db)

	odontologoRepository := odontologo.NewOdontologoRepository(storage)
	odontologoService := odontologo.NewOdontologoService(odontologoRepository)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService)

	pacienteRepository := paciente.NewPacienteRepository(storage)
	pacienteService := paciente.NewPacienteService(pacienteRepository)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	turnoRepository := turno.NewTurnoRepository(storage)
	turnoService := turno.NewTurnoService(turnoRepository)
	turnoHandler := handler.NewTurnoHandler(turnoService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	odontologos := r.Group("/odontologos")

	{
		odontologos.GET(":id", odontologoHandler.GetByID())
		odontologos.POST("", middleware.Authentication(), odontologoHandler.Post())
		odontologos.PUT(":id", middleware.Authentication(), odontologoHandler.Put())
		odontologos.PATCH(":id", middleware.Authentication(), odontologoHandler.Patch())
		odontologos.DELETE(":id", middleware.Authentication(), odontologoHandler.Delete())
	}

	pacientes := r.Group("/pacientes")

	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())
	}

	turnos := r.Group("/turnos")

	{
		turnos.POST("", middleware.Authentication(), turnoHandler.Post())
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.PATCH(":id", middleware.Authentication(), turnoHandler.Patch())
		turnos.PUT(":id", middleware.Authentication(), turnoHandler.Put())
		turnos.DELETE(":id", middleware.Authentication(), turnoHandler.Delete())
		turnos.GET("", turnoHandler.GetByDNI())
		turnos.POST("/DniMat", middleware.Authentication(), turnoHandler.PostDniMat())
	}

	r.Run(":8080")

}
