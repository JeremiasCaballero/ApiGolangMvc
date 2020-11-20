package vuelomodel

import (
	"database/sql"
	"fmt"
	// comment
)

// Database ..
type Database struct {
	DB *sql.DB
}

// Vuelo ..
type Vuelo struct {
	ID      int
	Destino string
}

// NewDataBase ..
func NewDataBase() Database {
	db, err := sql.Open("mysql", "root:@/vuelo")

	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Conection suceful")
	return Database{db}
}

// AddVuelo ..
func (db *Database) AddVuelo(d string) {
	insert, err := db.DB.Query("INSERT INTO vuelo VALUES(null,?)", d)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
