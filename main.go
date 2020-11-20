package main

import (
	vuelomodel "github.com/JeremiasCaballero/ApiRestGolang/models/vuelo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := vuelomodel.NewDataBase()
	defer db.DB.Close()
	db.AddVuelo("Maimi")

}
