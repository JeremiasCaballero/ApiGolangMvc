package vuelocontroller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

//VueloController ..
type VueloController struct {
	Model *sql.DB
}

// Vuelo ..
type Vuelo struct {
	ID      int    `json:"id"`
	Destino string `json:"Vuelo"`
}

// NewController ..
func NewController(db *sql.DB) {
	// recibe una instancia de la db para que las demas funciones la usen

}

// AddVuelo ..
func AddVuelo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Esta andando llegaste hasta aca",
	})
}
