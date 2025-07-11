package fuenteLectura

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	fuente_lectura_model "main/source/modules/fuenteLectura/model"
)

var Service = base_service.NewService[base_service.Default[fuente_lectura_model.FuenteLecturaStruct]](*fuente_lectura_model.Model)

func Init() {
	print("FuenteLectura Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/fuenteLectura")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
