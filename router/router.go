package router

import (
	vuelocontroller "github.com/JeremiasCaballero/ApiRestGolang/Controllers/vueloconntroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.POST("/addvuelo", vuelocontroller.AddVuelo)
	r.GET("/getVuelo/:id", vuelocontroller.GetVuelo)
	r.GET("/getvuelos", vuelocontroller.GetVuelos)
	r.DELETE("/deletevuelo/:id", vuelocontroller.DeleteVuelo)
	r.PUT("/updatevuelo/:id", vuelocontroller.UpdateVuelo)
	r.Run()

}
