package main

import (
	vuelocontroller "github.com/JeremiasCaballero/ApiRestGolang/Controllers/vueloconntroller"
	vuelomodel "github.com/JeremiasCaballero/ApiRestGolang/models/vuelo"
	"github.com/JeremiasCaballero/ApiRestGolang/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := vuelomodel.NewDataBase()
	defer db.DB.Close()
	// le paso la instancia de la base de datos
	vuelocontroller.NewController(&db)
	router.Start()

}
