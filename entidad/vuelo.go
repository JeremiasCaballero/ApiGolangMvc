package entidad

// VueloJSON ..
type VueloJSON struct {
	Destino string `json:"destino" binding:"required"`
}
