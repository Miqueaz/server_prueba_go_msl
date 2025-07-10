package modulos

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	modulos_model "main/source/modules/modulos/model"
)

var Service = base_service.NewService[base_service.Default[modulos_model.ModulosStruct]](*modulos_model.Model)

func Init() {
	print("Permisos Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/modulos")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
