package vuelomodel

import (
	"database/sql"
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
		panic("Failed to connect to database!")
	}
	err = db.Ping()
	if err != nil {
		panic("error al conectar")
	}
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

// DeleteVuelo ..
func (db *Database) DeleteVuelo(id int) {
	delete, err := db.DB.Query("DELETE FROM vuelo WHERE id_vuelo = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}
