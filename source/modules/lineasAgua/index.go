package lineasAgua

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	lineas_agua_model "main/source/modules/lineasAgua/model"
)

var Service = base_service.NewService[base_service.Default[lineas_agua_model.LineasAguaStruct]](*lineas_agua_model.Model)

func Init() {
	print("LineasAgua Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/lineasAgua")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
