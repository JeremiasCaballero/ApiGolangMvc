package vuelocontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeremiasCaballero/ApiRestGolang/entidad"
	vuelomodel "github.com/JeremiasCaballero/ApiRestGolang/models/vuelo"
	"github.com/gin-gonic/gin"
)

var modelo *vuelomodel.Database

// NewController ..
func NewController(model *vuelomodel.Database) {
	modelo = model
}

// AddVuelo ..
func AddVuelo(ctx *gin.Context) {
	var param entidad.VueloJSON

	ctx.BindJSON(&param)
	modelo.AddVuelo(param.Destino)

	ctx.JSON(http.StatusCreated, gin.H{
		"Objeto Creado": &param,
	})
}

// DeleteVuelo ..
func DeleteVuelo(ctx *gin.Context) {
	// controlar que sea barra id
	id := ctx.Param("id")

	var convertido int

	convertido, err := strconv.Atoi(id)
	if err != nil {
		panic("No es un integer")
	}

	fmt.Println(convertido)

	modelo.DeleteVuelo(convertido)
	ctx.JSON(http.StatusCreated, gin.H{
		"Objeto borrado": convertido,
	})
}

// GetVuelos ..
func GetVuelos(ctx *gin.Context) {
	v := modelo.GetVuelos()
	ctx.JSON(http.StatusCreated, gin.H{
		"Object": v,
	})
}

// GetVuelo ..
func GetVuelo(ctx *gin.Context) {
	id := ctx.Param("id")

	var convertido int

	convertido, err := strconv.Atoi(id)

	if err != nil {
		panic("No es un integer")
	}
	var param entidad.VueloJSON

	ctx.BindJSON(&param)

	modelo.getVuelo(param)
	ctx.JSON(http.StatusOK, gin.H{
		"Objeto Creado": "adada",
	})
}

// UpdateVuelo ..
func UpdateVuelo(ctx *gin.Context) {
	id := ctx.Param("id")

	var convertido int

	convertido, err := strconv.Atoi(id)

	if err != nil {
		panic("No es un integer")
	}
	var param entidad.VueloJSON
	ctx.BindJSON(&param)
	modelo.UpdateVuelo(convertido, param.Destino)

	ctx.JSON(http.StatusOK, gin.H{
		"Objeto editado": &param,
	})
}
