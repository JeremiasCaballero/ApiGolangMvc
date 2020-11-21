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
func (db *Database) getVuelo(id int) *Vuelo{
	var v *Vuelo
	select, err := db.DB.Query("SELECT *FROM vuelo WHERE id_vuelo = ?", id)
	if err != nil {
		panic(err.Error())
	}
	v = append(select)
	defer select.close()
	return select
}
// GetVuelos ..
func (db *Database) GetVuelos() []*Vuelo {
	var v []*Vuelo

	s, err := db.DB.Query("SELECT *FROM vuelo")

	if err != nil {
		panic(err.Error())
	}

	for s.Next() {
		var e Vuelo
		s.Scan(&e.ID, &e.Destino)
		v = append(v, &e)

	}
	defer s.Close()
	return v
}

// UpdateVuelo ..
func (db *Database) UpdateVuelo(id int, d string) {
	update, err := db.DB.Query("UPDATE  vuelo SET destino = ? WHERE id_vuelo = ?", d, id)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}
