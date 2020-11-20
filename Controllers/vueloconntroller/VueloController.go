package vuelocontroller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//VueloController ..
type VueloController struct {
	C *Vuelo
}

// Vuelo ..
type Vuelo struct {
	ID      int    `json:"id"`
	Destino string `json:"Vuelo"`
}

// NewController ..
func NewController() VueloController {
	router := gin.Default()
	router.GET("/addvuelo", AddVuelo)

	return VueloController{C}
}

func (v *VueloController) AddVuelo(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// panic
	}

	var vuelo Vuelo
	if err = json.Unmarshal(body, &car); err != nil {
		panic(error.Error())
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "car added",
	})
}
