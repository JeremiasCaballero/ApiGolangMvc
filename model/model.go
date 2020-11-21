package model

import "gorm.io/gorm"

// Vuelo es la entidad que se va a persistir en la base de datos
type Vuelo struct {
	gorm.Model
	ID      int    `gorm:"column:id;type:int"`
	Destino string `gorm:"column:destino;type:text(255)"`
}
