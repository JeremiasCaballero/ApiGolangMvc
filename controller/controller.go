package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jeremiascaballero/ApiGolangMvc/model"
	"github.com/jeremiascaballero/ApiGolangMvc/service"
)

// HTTPController ...
type HTTPController struct {
	router  *gin.Engine
	service *service.Vuelo
}

// NewHTTPController ...
func NewHTTPController(s *service.Vuelo) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // registro los endpoints
	return HTTPController{r, s}
}

// Run ejecuta el controller
func (c *HTTPController) Run() {
	c.router.Run()
}

func makeEndpoints(r *gin.Engine, s *service.Vuelo) {
	r.GET("/vuelos/:id", makeFindHandler(s))
	r.POST("/addvuelo", makeAddHandler(s))
	r.DELETE("/deletevuelo/:id", makeDeleteHandler(s))
}

func makeFindHandler(s *service.Vuelo) gin.HandlerFunc {
	// Fijate que aca devuelvo una funcion y no un valor
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		v := (*s).FindByID(uint(i))

		fmt.Println(&v)
		c.JSON(http.StatusOK, gin.H{
			"vuelo": v,
		})
	}
}

func makeAddHandler(s *service.Vuelo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param *model.Vuelo

		c.BindJSON(&param)

		(*s).Add(param)

		c.JSON(http.StatusOK, gin.H{
			"vuelo": &param,
		})
	}
}

func makeDeleteHandler(S *service.Vuelo) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		c.BindJSON(i)

		(*S).Remove(uint(i))

		c.JSON(http.StatusNoContent, gin.H{})
	}
}
