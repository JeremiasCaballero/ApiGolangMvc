package database

import (
	"github.com/jeremiascaballero/ApiGolangMvc/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	db *gorm.DB
}

// NewDatabase ...
func NewDatabase() *Database { //dsn := ""

	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/vuelo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Database{db}
}

// FindAll devuelve todos los vuelos registrados
func (database *Database) FindAll() []model.Vuelo {
	var vuelos []model.Vuelo
	database.db.Find(&vuelos)
	return vuelos
}

// CreateSchema create el esquema a partir de las entidades.
func (database *Database) CreateSchema() {
	database.db.AutoMigrate(&model.Vuelo{})
}

// Add inserta un vuelo
func (database *Database) Add(v *model.Vuelo) {
	database.db.Create(v)
}

// FindByID ...
func (database *Database) FindByID(ID uint) *model.Vuelo {
	var v model.Vuelo
	database.db.First(&v, ID)
	return &v
}

// Remove ...
func (database *Database) Remove(i uint) {
	var v model.Vuelo
	database.db.Unscoped().Delete(&v, i)
}
