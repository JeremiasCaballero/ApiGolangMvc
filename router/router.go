package router

import (
	vuelocontroller "github.com/JeremiasCaballero/ApiRestGolang/Controllers/vueloconntroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.POST("/addvuelo", vuelocontroller.AddVuelo)
	r.GET("/getVuelo", vuelocontroller.GetVuelo)
	r.DELETE("/deletevuelo/:id", vuelocontroller.DeleteVuelo)
	r.Run()

}
