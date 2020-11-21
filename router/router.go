package router

import (
	vuelocontroller "github.com/JeremiasCaballero/ApiRestGolang/Controllers/vueloconntroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.GET("/prueba", vuelocontroller.AddVuelo)

	r.Run()

}
