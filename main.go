package main

import (
	"github.com/jeremiascaballero/ApiGolangMvc/controller"
	"github.com/jeremiascaballero/ApiGolangMvc/model/database"
	"github.com/jeremiascaballero/ApiGolangMvc/service"
)

func main() {
	// Creo una instancia de la base de datos
	db := database.NewDatabase() // aca esta hardcoded el nombre de la base de datos
	// Creo el esquema
	db.CreateSchema()

	// creo el servicio y le paso la base de datos
	service := service.NewInstance(db)

	http := controller.NewHTTPController(&service)
	http.Run()
}
