package service

import (
	"github.com/jeremiascaballero/ApiGolangMvc/model"
	"github.com/jeremiascaballero/ApiGolangMvc/model/database"
)

// Vuelo define el comportamiento del servicio de vuelos
type Vuelo interface {
	Add(*model.Vuelo) *model.Vuelo
	FindByID(ID uint) *model.Vuelo
	Remove(ID uint)
}

type service struct {
	db *database.Database
}

// NewInstance devuelve una instancia del servicio de vuelos
func NewInstance(db *database.Database) Vuelo {
	return &service{db}
}

func (s *service) Add(v *model.Vuelo) *model.Vuelo {
	s.db.Add(v)
	return v
}

func (s *service) FindByID(ID uint) *model.Vuelo {
	return s.db.FindByID(ID)
}

func (s *service) Remove(ID uint) {
}
